package managers

import (
	"errors"
	"fmt"
	"math"
	"paper-tracker/config"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

	"github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
)

var learningManager LearningManager

type LearningManager interface {
	StartLearning(trackerID models.TrackerID) (learnTimeSec int, err error)
	FinishLearning(trackerID models.TrackerID, roomID models.RoomID, ssids []string) error
	GetLearningStatus(trackerID models.TrackerID) (done bool, ssids []string, err error)
	CancelLearning(trackerID models.TrackerID) error
	NewLearningTrackingData(trackerID models.TrackerID, scanRes []*models.ScanResult) error
}

type LearningManagerImpl struct {
	scanResultRep repositories.ScanResultRepository
}

func CreateLearningManager(scanResultRep repositories.ScanResultRepository) LearningManager {
	if learningManager != nil {
		return learningManager
	}

	learningManager = &LearningManagerImpl{
		scanResultRep: scanResultRep,
	}

	return learningManager
}

func GetLearningManager() LearningManager {
	return learningManager
}

func (mgr *LearningManagerImpl) StartLearning(trackerID models.TrackerID) (learnTimeSec int, err error) {
	learnLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		return
	} else if tracker.Status != models.TrackerStatusIdle {
		err = errors.New("Given tracker is not in idle mode")
		return
	}

	err = mgr.scanResultRep.DeleteForTracker(trackerID)
	if err != nil {
		learnLog.WithError(err).Error("Failed to delete scan results for tracker")
		return
	}

	go mgr.learningRoutine(trackerID, learnLog)

	// Send a learn time 20% higher to take response time of tracker into account
	learnCount := config.GetInt(config.KeyCmdLearnCount)
	learnSleepSec := config.GetInt(config.KeyCmdLearnSleep)
	learnTime := float64(learnCount) * float64(learnSleepSec) * 1.2
	learnTimeSec = int(math.RoundToEven(learnTime))
	return
}

func (mgr *LearningManagerImpl) learningRoutine(trackerID models.TrackerID, logger *log.Entry) {
	defer ginkgo.GinkgoRecover() //FIXME: Leave this in for now as sometimes the unit tests crash in this goroutine

	logger.Info("Start learning routine")

	logger.Info("Set tracker status to learning")
	err := GetTrackerManager().SetTrackerStatus(trackerID, models.TrackerStatusLearning)
	if err != nil {
		return
	}

	canceled := mgr.checkLearningCanceled(trackerID, logger)

	if !canceled {
		logger.Info("Set tracker status to learning finished")
		err = GetTrackerManager().SetTrackerStatus(trackerID, models.TrackerStatusLearningFinished)
		if err != nil {
			return
		}
	} else {
		logger.Info("Tracker learning got canceled")
	}
}

// Returns whether the learning was canceled and also "waits" the learning time
func (mgr *LearningManagerImpl) checkLearningCanceled(trackerID models.TrackerID, logger *log.Entry) bool {
	logger.Info("Start checking for canceled learning")

	learnCount := config.GetInt(config.KeyCmdLearnCount)
	learnSleepSec := config.GetInt(config.KeyCmdLearnSleep)
	for it := 0; it < learnCount; it++ {
		tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
		if err == nil && tracker.Status != models.TrackerStatusLearning {
			return true
		}

		time.Sleep(time.Duration(learnSleepSec) * time.Second)
	}
	logger.Info("Finished checking for canceled learning")

	return false
}

func (mgr *LearningManagerImpl) NewLearningTrackingData(trackerID models.TrackerID, scanRes []*models.ScanResult) (err error) {
	for _, scan := range scanRes {
		scan.TrackerID = trackerID
	}

	err = mgr.scanResultRep.CreateAll(scanRes)
	return
}

func (mgr *LearningManagerImpl) FinishLearning(trackerID models.TrackerID, roomID models.RoomID, ssids []string) (err error) {
	finishLearningLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		finishLearningLog.WithError(err).Error("Failed to get tracker")
		err = fmt.Errorf("tracker: %v", err)
		return
	}

	if tracker.Status != models.TrackerStatusLearningFinished {
		err = errors.New("Tracker is not in status LearningFinished")
		return
	}

	room, err := GetRoomManager().GetRoomByID(roomID)
	if err != nil {
		finishLearningLog.WithError(err).Error("Failed to get room")
		err = fmt.Errorf("room: %v", err)
		return
	}

	scanResults, err := mgr.scanResultRep.GetAllForTracker(tracker.ID)
	if err != nil {
		finishLearningLog.WithError(err).Error("Failed to get the new scan results")
		err = fmt.Errorf("scanResults: %v", err)
		return
	}

	trackingData := GetTrackingManager().ConsolidateScanResults(scanResults)
	room.TrackingData = trackingData
	room.IsLearned = true
	err = GetRoomManager().UpdateRoom(room)
	if err != nil {
		finishLearningLog.WithError(err).Error("Could not save room")
		err = fmt.Errorf("room: %v", err)
	}

	err = GetTrackerManager().SetTrackerStatus(tracker.ID, models.TrackerStatusIdle)
	if err != nil {
		finishLearningLog.WithError(err).Error("Failed to set tracker status to idle after learning")
		return
	}

	return
}

func (mgr *LearningManagerImpl) GetLearningStatus(trackerID models.TrackerID) (done bool, ssids []string, err error) {
	learningStatusLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		learningStatusLog.WithError(err).Error("Failed to get tracker")
		return
	} else if tracker.Status != models.TrackerStatusLearning && tracker.Status != models.TrackerStatusLearningFinished {
		err = errors.New("Tracker currently not in learning or learning finished status")
		return
	}

	done = tracker.Status == models.TrackerStatusLearningFinished

	scanRes, err := mgr.scanResultRep.GetAllForTracker(trackerID)
	if err != nil && !mgr.scanResultRep.IsRecordNotFoundError(err) {
		learningStatusLog.WithError(err).Error("Failed to get scan results for tracker")
		return
	} else if mgr.scanResultRep.IsRecordNotFoundError(err) {
		ssids = make([]string, 0)
		return
	}

	// Filter out duplicates through map and assemble slice
	ssidMap := make(map[string]bool, 0)
	for _, scan := range scanRes {
		ssidMap[scan.SSID] = true
	}
	ssids = make([]string, len(ssidMap))
	it := 0
	for ssid := range ssidMap {
		ssids[it] = ssid
		it++
	}

	return
}

func (mgr *LearningManagerImpl) CancelLearning(trackerID models.TrackerID) (err error) {
	cancelLearningLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		cancelLearningLog.WithError(err).Error("Failed to get tracker")
		return
	}

	if tracker.Status != models.TrackerStatusLearning && tracker.Status != models.TrackerStatusLearningFinished {
		cancelLearningLog.Error("Tracker not in learning status")
		return errors.New("Tracker not in learning status")
	}

	err = GetTrackerManager().SetTrackerStatus(trackerID, models.TrackerStatusIdle)
	if err != nil {
		cancelLearningLog.WithError(err).Error("Failed to set tracker status to idle while canceling learning - ignore for now")
		err = nil
	}
	return
}
