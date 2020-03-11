package managers

import (
	"errors"
	"fmt"
	"paper-tracker/models"
	"paper-tracker/models/communication"
	"paper-tracker/repositories"
	"sync"

	log "github.com/sirupsen/logrus"
)

var trackerManager *TrackerManager

type TrackerManager struct {
	trackerRep            repositories.TrackerRepository
	scanResultsCacheMutex sync.Mutex
	scanResultsCache      map[models.TrackerID][]*models.ScanResult
	idleSleepSec          int
	trackingSleepSec      int
	learningSleepSec      int
	done                  chan struct{}
}

func CreateTrackerManager(trackerRep repositories.TrackerRepository, idleSleepSec, trackingSleepSec, learningSleepSec int) *TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManager{
		trackerRep:       trackerRep,
		idleSleepSec:     idleSleepSec,
		trackingSleepSec: trackingSleepSec,
		learningSleepSec: learningSleepSec,
		done:             make(chan struct{}),
		scanResultsCache: make(map[models.TrackerID][]*models.ScanResult),
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
		cmd = &models.Command{
			Type:         models.CmdSleep,
			SleepTimeSec: mgr.idleSleepSec,
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

	return
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
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithError(err).Error("Failed to update tracker")
		return
	}
	return
}

func (mgr *TrackerManager) UpdateRoom(trackerID models.TrackerID, roomID models.RoomID) (err error) {
	updateLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		updateLog.WithField("err", err).Error("Failed to get tracker with id")
		return
	}

	tracker.LastRoom = roomID
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		updateLog.WithField("err", err).Error("Failed to update tracker")
		return
	}
	return
}

func (mgr *TrackerManager) NewTrackingData(trackerID models.TrackerID, isLastBatch bool, scanRes []*models.ScanResult) (err error) {
	trackingDataLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		trackingDataLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	if tracker.Status == models.TrackerStatusIdle || tracker.Status == models.TrackerStatusTracking {
		mgr.scanResultsCacheMutex.Lock()
		defer mgr.scanResultsCacheMutex.Unlock()
		mgr.scanResultsCache[tracker.ID] = append(mgr.scanResultsCache[tracker.ID], scanRes...)
		if isLastBatch {
			scanResults := mgr.scanResultsCache[tracker.ID]
			mgr.scanResultsCache[tracker.ID] = nil
			err = setMatchingRoomForTracker(tracker, scanResults)
			if err != nil {
				return
			}
		}

		err = setMatchingRoomForTracker(tracker, scanRes)
		if err != nil {
			log.Error(err)
			return
		}
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		err = errors.New("No tracking data expected")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unexpected tracking data")
	case models.TrackerStatusLearning:
		err = GetLearningManager().newLearningTrackingData(trackerID, scanRes)
	case models.TrackerStatusTracking:
		if isLastBatch {
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
	return GetTrackerManager().UpdateRoom(tracker.ID, bestMatch.ID)
}
