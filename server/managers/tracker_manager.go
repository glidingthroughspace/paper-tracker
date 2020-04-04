package managers

import (
	"errors"
	"fmt"
	"paper-tracker/models"
	"paper-tracker/models/communication"
	"paper-tracker/repositories"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var trackerManager *TrackerManager

type CachedScanResults struct {
	ResultID        uint64
	BatchesExpected uint8
	BatchesReceived uint8
	ScanResults     []*models.ScanResult
}

type TrackerManager struct {
	trackerRep            repositories.TrackerRepository
	scanResultsCacheMutex sync.Mutex
	scanResultsCache      map[models.TrackerID]CachedScanResults
	idleSleepSec          int
	trackingSleepSec      int
	learningSleepSec      int
	sendInfoSleepSec      int
	sendInfoIntervalSec   int
	workStartHour         int
	workEndHour           int
	workOnWeekend         bool
	done                  chan struct{}
}

func CreateTrackerManager(
	trackerRep repositories.TrackerRepository,
	idleSleepSec,
	trackingSleepSec,
	learningSleepSec,
	sendInfoSleepSec,
	sendInfoIntervalSec,
	workStartHour,
	workEndHour int,
	workOnWeekend bool) *TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManager{
		trackerRep:          trackerRep,
		idleSleepSec:        idleSleepSec,
		trackingSleepSec:    trackingSleepSec,
		learningSleepSec:    learningSleepSec,
		sendInfoSleepSec:    sendInfoSleepSec,
		sendInfoIntervalSec: sendInfoIntervalSec,
		done:                make(chan struct{}),
		scanResultsCache:    make(map[models.TrackerID]CachedScanResults),
	}

	return trackerManager
}

func GetTrackerManager() *TrackerManager {
	return trackerManager
}

func (mgr *TrackerManager) GetTrackerByID(trackerID models.TrackerID) (tracker *models.Tracker, err error) {
	tracker, err = mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		log.WithFields(log.Fields{"trackerID": trackerID, "err": err}).Error("Tracker not found")
		return
	}
	return
}

func (mgr *TrackerManager) GetAllTrackers() (trackers []*models.Tracker, err error) {
	trackers, err = mgr.trackerRep.GetAll()
	if err != nil {
		log.WithError(err).Error("Failed to get all trackers")
		return
	}
	return
}

func (mgr *TrackerManager) SetTrackerStatus(trackerID models.TrackerID, status models.TrackerStatus) (err error) {
	err = mgr.trackerRep.SetStatusByID(trackerID, status)
	if err != nil {
		log.WithFields(log.Fields{"trackerID": trackerID, "status": status, "err": err}).Error("Failed to set status of tracker")
		return
	}
	return
}

func (mgr *TrackerManager) UpdateTrackerLabel(trackerID models.TrackerID, label string) (tracker *models.Tracker, err error) {
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

func (mgr *TrackerManager) DeleteTracker(trackerID models.TrackerID) (err error) {
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

func (mgr *TrackerManager) NotifyNewTracker() (tracker *models.Tracker, err error) {
	tracker = &models.Tracker{Label: "New Tracker", Status: models.TrackerStatusIdle}
	err = mgr.trackerRep.Create(tracker)
	if err != nil {
		log.WithError(err).Error("Failed to create new tracker")
		return
	}
	return
}

func (mgr *TrackerManager) PollCommand(trackerID models.TrackerID) (cmd *models.Command, err error) {
	pollLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		pollLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		// If the tracker is idling, we want to periodically check for battery stats.
		if int(time.Since(tracker.LastBatteryUpdate).Seconds()) > mgr.sendInfoIntervalSec {
			cmd = &models.Command{
				Type:         models.CmdSendInformation,
				SleepTimeSec: mgr.sendInfoSleepSec,
			}
		} else {
			cmd = &models.Command{
				Type:         models.CmdSleep,
				SleepTimeSec: mgr.idleSleepSec,
			}
		}
	case models.TrackerStatusTracking:
		cmd = &models.Command{
			Type:         models.CmdSendTrackingInformation,
			SleepTimeSec: mgr.trackingSleepSec,
		}
	case models.TrackerStatusLearning:
		cmd = &models.Command{
			Type:         models.CmdSendTrackingInformation,
			SleepTimeSec: mgr.learningSleepSec,
		}
	}

	if !mgr.InWorkingHours() {
		cmd.SleepTimeSec = 100 //TODO: Figure out max sleep time
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
func (mgr *TrackerManager) InWorkingHours() bool {
	if mgr.workStartHour < 0 || mgr.workEndHour < 0 {
		return true
	}

	currentTime := time.Now().Local()
	if day := currentTime.Weekday(); !mgr.workOnWeekend && (day == time.Saturday || day == time.Sunday) {
		return false
	} else if hour := currentTime.Hour(); hour < mgr.workStartHour || hour > mgr.workEndHour {
		return false
	}
	return true
}

func (mgr *TrackerManager) UpdateFromResponse(trackerID models.TrackerID, resp communication.TrackerCmdResponse) (err error) {
	updateLog := log.WithFields(log.Fields{"trackerID": trackerID, "resp": resp})

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		updateLog.WithError(err).Error("Failed to get tracker with id")
		return
	}

	tracker.BatteryPercentage = resp.BatteryPercentage
	tracker.IsCharging = resp.IsCharging
	tracker.LastBatteryUpdate = time.Now()

	updateLog.Debugf("Tracker %ds battery is at %d%% capacity", tracker.ID, resp.BatteryPercentage)
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithError(err).Error("Failed to update tracker")
		return
	}
	return
}

func (mgr *TrackerManager) UpdateRoom(tracker *models.Tracker, roomID models.RoomID) (err error) {
	updateLog := log.WithFields(log.Fields{"trackerID": tracker.ID, "roomID": roomID})

	tracker.LastRoom = roomID
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithField("err", err).Error("Failed to update tracker")
		return
	}
	return
}

func (mgr *TrackerManager) NewTrackingData(trackerID models.TrackerID, resultID uint64, batchCount uint8, scanRes []*models.ScanResult) (err error) {
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
			err = setMatchingRoomForTracker(tracker, scanResults)
			if err != nil {
				return
			}
			log.Debugf("Last known room ID for tracker is: %v", tracker.LastRoom)
		}
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		err = errors.New("No tracking data expected")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unexpected tracking data")
	case models.TrackerStatusLearning:
		err = GetLearningManager().newLearningTrackingData(trackerID, scanRes)
	case models.TrackerStatusTracking:
		if mgr.receivedAllBatchesForTracker(tracker.ID) {
			err = GetWorkflowExecManager().ProgressToTrackerRoom(tracker.ID, tracker.LastRoom)
		}
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

func (mgr *TrackerManager) receivedAllBatchesForTracker(trackerID models.TrackerID) bool {
	return mgr.scanResultsCache[trackerID].BatchesReceived == mgr.scanResultsCache[trackerID].BatchesExpected
}

func setMatchingRoomForTracker(tracker *models.Tracker, scanResults []*models.ScanResult) error {
	rooms, err := GetRoomManager().GetAllRooms()
	if err != nil {
		err = fmt.Errorf("could not get rooms: %w", err)
		return err
	}
	bestMatch := GetTrackingManager().GetRoomMatchingBest(rooms, scanResults)
	if bestMatch == nil {
		err = fmt.Errorf("no matching room found")
		return err
	}
	return GetTrackerManager().UpdateRoom(tracker, bestMatch.ID)
}
