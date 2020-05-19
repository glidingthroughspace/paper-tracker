package managers

import (
	"errors"
	"fmt"
	"paper-tracker/config"
	"paper-tracker/models"
	"paper-tracker/models/communication"
	"paper-tracker/repositories"
	"paper-tracker/utils"
	"sync"
	"time"

	"github.com/jinzhu/now"
	log "github.com/sirupsen/logrus"
)

var trackerManager TrackerManager

type TrackerManager interface {
	GetTrackerByID(trackerID models.TrackerID) (*models.Tracker, error)
	GetAllTrackers() ([]*models.Tracker, error)
	SetTrackerStatus(trackerID models.TrackerID, status models.TrackerStatus) error
	UpdateTrackerLabel(trackerID models.TrackerID, label string) (*models.Tracker, error)
	DeleteTracker(trackerID models.TrackerID) error
	NotifyNewTracker() (*models.Tracker, error)
	PollCommand(trackerID models.TrackerID) (*models.Command, error)
	InWorkingHours(currentTime time.Time) (inHours bool, toStart time.Duration)
	UpdateFromResponse(trackerID models.TrackerID, resp communication.TrackerCmdResponse) error
	UpdateRoom(tracker *models.Tracker, roomID models.RoomID) error
	NewTrackingData(trackerID models.TrackerID, resultID uint64, batchCount uint8, scanRes []*models.ScanResult) error
}

type CachedScanResults struct {
	ResultID        uint64
	BatchesExpected uint8
	BatchesReceived uint8
	ScanResults     []*models.ScanResult
}

type TrackerManagerImpl struct {
	trackerRep            repositories.TrackerRepository
	scanResultsCacheMutex sync.Mutex
	scanResultsCache      map[models.TrackerID]CachedScanResults
	trackingCounts        map[models.TrackerID][]map[models.RoomID]float64
	done                  chan struct{}
}

func CreateTrackerManager(trackerRep repositories.TrackerRepository) TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManagerImpl{
		trackerRep:       trackerRep,
		done:             make(chan struct{}),
		scanResultsCache: make(map[models.TrackerID]CachedScanResults),
		trackingCounts:   make(map[models.TrackerID][]map[models.RoomID]float64),
	}

	return trackerManager
}

func GetTrackerManager() TrackerManager {
	return trackerManager
}

func (mgr *TrackerManagerImpl) GetTrackerByID(trackerID models.TrackerID) (tracker *models.Tracker, err error) {
	tracker, err = mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		log.WithFields(log.Fields{"trackerID": trackerID, "err": err}).Error("Tracker not found")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) GetAllTrackers() (trackers []*models.Tracker, err error) {
	trackers, err = mgr.trackerRep.GetAll()
	if err != nil {
		log.WithError(err).Error("Failed to get all trackers")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) SetTrackerStatus(trackerID models.TrackerID, status models.TrackerStatus) (err error) {
	err = mgr.trackerRep.SetStatusByID(trackerID, status)
	if err != nil {
		log.WithFields(log.Fields{"trackerID": trackerID, "status": status, "err": err}).Error("Failed to set status of tracker")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) UpdateTrackerLabel(trackerID models.TrackerID, label string) (tracker *models.Tracker, err error) {
	setLabelLog := log.WithFields(log.Fields{"trackerID": trackerID, "label": label})

	tracker, err = mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		setLabelLog.WithError(err).Error("Failed to get tracker")
		return
	}

	tracker.Label = label
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		log.WithError(err).Error("Failed to set label of tracker")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) DeleteTracker(trackerID models.TrackerID) (err error) {
	deleteLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.GetTrackerByID(trackerID)
	if err != nil {
		deleteLog.WithError(err).Error("Failed to get tracker to delete")
		return
	}

	if tracker.Status != models.TrackerStatusIdle {
		deleteLog.Error("Can not delete tracker that is not in idle")
		return errors.New("Can not delete tracker that is not in idle")
	}

	err = mgr.trackerRep.Delete(trackerID)
	if err != nil {
		deleteLog.WithError(err).Error("Failed to delete tracker")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) NotifyNewTracker() (tracker *models.Tracker, err error) {
	tracker = &models.Tracker{Label: "New Tracker", Status: models.TrackerStatusIdle}
	err = mgr.trackerRep.Create(tracker)
	if err != nil {
		log.WithError(err).Error("Failed to create new tracker")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) PollCommand(trackerID models.TrackerID) (cmd *models.Command, err error) {
	pollLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		pollLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		// If the tracker is idling, we want to periodically check for battery stats.
		if int(time.Since(tracker.LastBatteryUpdate).Seconds()) > config.GetInt(config.KeyCmdInfoInterval) {
			cmd = &models.Command{
				Type:         models.CmdSendInformation,
				SleepTimeSec: config.GetInt(config.KeyCmdInfoSleep),
			}
		} else {
			cmd = &models.Command{
				Type:         models.CmdSleep,
				SleepTimeSec: config.GetInt(config.KeyCmdIdleSleep),
			}
		}
	case models.TrackerStatusTracking:
		sleepTime := 1
		// FIXME: Make this configurable
		if len(mgr.trackingCounts[tracker.ID]) >= 2 {
			sleepTime = config.GetInt(config.KeyCmdTrackSleep)
		}
		cmd = &models.Command{
			Type:         models.CmdSendTrackingInformation,
			SleepTimeSec: sleepTime,
		}
	case models.TrackerStatusLearning:
		cmd = &models.Command{
			Type:         models.CmdSendTrackingInformation,
			SleepTimeSec: config.GetInt(config.KeyCmdLearnSleep),
		}
	}

	if inWorkHours, toWorkHours := mgr.InWorkingHours(time.Now().Local()); !inWorkHours {
		maxSleepSec := config.GetInt(config.KeyCmdMaxSleep)
		toWorkHoursSec := int(toWorkHours.Seconds())
		if toWorkHoursSec > maxSleepSec {
			cmd.SleepTimeSec = maxSleepSec
		} else {
			cmd.SleepTimeSec = toWorkHoursSec
		}
	}

	tracker.LastPoll = time.Now()
	tracker.LastSleepTimeSec = cmd.SleepTimeSec
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		pollLog.WithError(err).Error("Failed to update last poll time of tracker, ignoring")
		err = nil
	}

	return
}

//TODO: Write test and somehow mock time.Now() for that
// InWorkingHours returns whether we are currently in working hours and if not, how long it will be until working hour
func (mgr *TrackerManagerImpl) InWorkingHours(currentTime time.Time) (inHours bool, toStart time.Duration) {
	workStartHour := config.GetInt(config.KeyWorkStartHour)
	workEndHour := config.GetInt(config.KeyWorkEndHour)
	workOnWeekend := config.GetBool(config.KeyWorkOnWeekend)

	return mgr.inWorkingHours(currentTime, workStartHour, workEndHour, workOnWeekend)
}

func (mgr *TrackerManagerImpl) inWorkingHours(currentTime time.Time, workStartHour, workEndHour int, workOnWeekend bool) (inHours bool, toStart time.Duration) {
	if workStartHour < 0 || workEndHour < 0 {
		return true, time.Duration(0)
	}

	toStartDuration := time.Duration(workStartHour) * time.Hour // Time from midnight to work start
	oneDay := time.Hour * 24

	workStartTime := todayWithSetHour(currentTime, workStartHour)
	workEndTime := todayWithSetHour(currentTime, workEndHour)

	var startTime time.Time
	if itIsWeekend(currentTime) && !workOnWeekend {
		// If we don't work on the weekend, check if we have Saturday/Sunday
		// startTime of next working day is the Monday of next week at the proper hour
		startTime = now.With(currentTime.Add(oneDay * 3)).Monday().Add(toStartDuration)
	} else if currentTime.Before(workStartTime) {
		// If we have a workday but are BEFORE the working hours
		// startTime is today at the proper hour
		startTime = now.With(currentTime.Add(oneDay)).BeginningOfDay().Add(toStartDuration)
	} else if currentTime.After(workEndTime) {
		// If we have a workday but are AFTER the working hours
		// startTime is the next day at the proper hour (ignore transition to the weekend for now)
		startTime = now.With(currentTime).BeginningOfDay().Add(toStartDuration)
	} else {
		inHours = true
	}
	toStart = startTime.Sub(currentTime)
	return
}

func todayWithSetHour(today time.Time, hour int) time.Time {
	return time.Date(today.Year(), today.Month(), today.Day(), hour, 0, 0, 0, today.Location())
}

func itIsWeekend(today time.Time) bool {
	return today.Weekday() == time.Saturday || today.Weekday() == time.Sunday
}

func (mgr *TrackerManagerImpl) UpdateFromResponse(trackerID models.TrackerID, resp communication.TrackerCmdResponse) (err error) {
	updateLog := log.WithFields(log.Fields{"trackerID": trackerID, "resp": resp})

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		updateLog.WithError(err).Error("Failed to get tracker with id")
		return
	}

	tracker.BatteryPercentage = resp.BatteryPercentage
	tracker.IsCharging = resp.IsCharging
	tracker.LastBatteryUpdate = time.Now()

	lowBatteryThreshold := config.GetInt(config.KeyLowBatteryThreshold)
	if tracker.BatteryPercentage < lowBatteryThreshold && !tracker.IsCharging && !tracker.LowBatteryNotified {
		err = utils.SendMail(
			fmt.Sprintf("Paper-Tracker: '%s' has a low battery", tracker.Label),
			fmt.Sprintf("The battery of the tracker '%s' only has %d%% left! Please charge the tracker as soon as possible.", tracker.Label, tracker.BatteryPercentage))
		if err != nil {
			updateLog.WithError(err).Warn("Failed to send low battery notification - ignore for now")
			err = nil
		} else {
			tracker.LowBatteryNotified = true
		}
	} else if tracker.IsCharging {
		tracker.LowBatteryNotified = false
	}

	updateLog.Debugf("Tracker %ds battery is at %d%% capacity", tracker.ID, resp.BatteryPercentage)
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithError(err).Error("Failed to update tracker")
		return
	}

	return
}

func (mgr *TrackerManagerImpl) UpdateRoom(tracker *models.Tracker, roomID models.RoomID) (err error) {
	updateLog := log.WithFields(log.Fields{"trackerID": tracker.ID, "roomID": roomID})

	tracker.LastRoom = roomID
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithField("err", err).Error("Failed to update tracker")
		return
	}
	return
}

func (mgr *TrackerManagerImpl) NewTrackingData(trackerID models.TrackerID, resultID uint64, batchCount uint8, scanRes []*models.ScanResult) (err error) {
	trackingDataLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		trackingDataLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	if tracker.Status == models.TrackerStatusIdle || tracker.Status == models.TrackerStatusTracking {
		mgr.scanResultsCacheMutex.Lock()
		defer mgr.scanResultsCacheMutex.Unlock()
		if mgr.scanResultsCache[tracker.ID].ResultID != resultID {
			trackingDataLog.Infof("Received scan results with result ID %v, replacing scan results for result ID %v", resultID, mgr.scanResultsCache[tracker.ID].ResultID)
			trackingDataLog.Debugf("Expecting %d scan result batches", batchCount)
			mgr.scanResultsCache[tracker.ID] = CachedScanResults{
				ResultID:        resultID,
				BatchesExpected: batchCount,
				BatchesReceived: 1,
				ScanResults:     scanRes,
			}
		} else {
			trackingDataLog.WithFields(log.Fields{"expected": mgr.scanResultsCache[tracker.ID].BatchesExpected, "received": mgr.scanResultsCache[tracker.ID].BatchesReceived + 1}).Debugf("Got additional scan result batch")
			mgr.scanResultsCache[tracker.ID] = CachedScanResults{
				ResultID:        resultID,
				BatchesExpected: batchCount,
				BatchesReceived: mgr.scanResultsCache[tracker.ID].BatchesReceived + 1,
				ScanResults:     append(mgr.scanResultsCache[tracker.ID].ScanResults, scanRes...),
			}
		}

		if mgr.receivedAllBatchesForTracker(tracker.ID) {
			scanResults := mgr.scanResultsCache[tracker.ID].ScanResults
			var scoredRooms map[models.RoomID]float64
			scoredRooms, err = scoreRoomsForTracker(tracker, scanResults)
			if err != nil {
				return
			}
			mgr.trackingCounts[tracker.ID] = append(mgr.trackingCounts[tracker.ID], scoredRooms)
			if len(mgr.trackingCounts[tracker.ID]) >= 3 {
				err = setMatchingRoomForTracker(tracker, mgr.trackingCounts[tracker.ID])
				mgr.trackingCounts[tracker.ID] = make([]map[models.RoomID]float64, 0, 0)
				if err != nil {
					return
				}
				log.Debugf("Last known room ID for tracker is: %v", tracker.LastRoom)
				err = GetWorkflowExecManager().ProgressToTrackerRoom(tracker.ID, tracker.LastRoom)
				if err != nil {
					return
				}
			}
		}
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		err = errors.New("No tracking data expected")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unexpected tracking data")
	case models.TrackerStatusLearning:
		err = GetLearningManager().NewLearningTrackingData(trackerID, scanRes)
	case models.TrackerStatusTracking:
		// empty
	default:
		err = errors.New("Unknown tracker status")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unknown tracker status")
	}

	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (mgr *TrackerManagerImpl) receivedAllBatchesForTracker(trackerID models.TrackerID) bool {
	return mgr.scanResultsCache[trackerID].BatchesReceived == mgr.scanResultsCache[trackerID].BatchesExpected
}

func scoreRoomsForTracker(tracker *models.Tracker, scanResults []*models.ScanResult) (map[models.RoomID]float64, error) {
	rooms, err := GetRoomManager().GetAllRooms()
	if err != nil {
		err = fmt.Errorf("could not get rooms: %w", err)
		return nil, err
	}
	log.Debug("Scored rooms")
	scoredRooms := GetTrackingManager().ScoreRoomsForScanResults(rooms, scanResults)
	return scoredRooms, nil
}

func setMatchingRoomForTracker(tracker *models.Tracker, scoredRooms []map[models.RoomID]float64) error {
	bestMatch := GetTrackingManager().GetRoomMatchingBest(scoredRooms)
	if bestMatch <= 0 {
		return fmt.Errorf("no matching room found")
	}
	return GetTrackerManager().UpdateRoom(tracker, bestMatch)
}
