package managers

import (
	"errors"
	"fmt"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

	"github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
)

var learningManager *LearningManager

type LearningManager struct {
	scanResultRep        repositories.ScanResultRepository
	learnCount           int
	sleepBetweenLearnSec int
}

func CreateLearningManager(scanResultRep repositories.ScanResultRepository, learnCount, sleepBetweenLearnSec int) *LearningManager {
	if learningManager != nil {
		return learningManager
	}

	learningManager = &LearningManager{
		scanResultRep:        scanResultRep,
		learnCount:           learnCount,
		sleepBetweenLearnSec: sleepBetweenLearnSec,
	}

	return learningManager
}

func GetLearningManager() *LearningManager {
	return learningManager
}

func (mgr *LearningManager) StartLearning(trackerID int) (learnTimeSec int, err error) {
	learnLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		return
	} else if tracker.Status != models.StatusIdle {
		err = errors.New("Given tracker is not in idle mode")
		return
	}

	err = mgr.scanResultRep.DeleteForTracker(trackerID)
	if err != nil {
		learnLog.WithField("err", err).Error("Failed to delete scan results for tracker")
		return
	}

	go mgr.learningRoutine(trackerID, learnLog)

	learnTimeSec = mgr.learnCount * mgr.sleepBetweenLearnSec
	return
}

func (mgr *LearningManager) learningRoutine(trackerID int, logger *log.Entry) {
	defer ginkgo.GinkgoRecover() //FIXME: Leave this in for now as sometimes the unit tests crash in this goroutine

	logger.Info("Start learning routine")

	logger.Info("Set tracker status to learning")
	err := GetTrackerManager().SetTrackerStatus(trackerID, models.StatusLearning)
	if err != nil {
		return
	}

	mgr.learningCreateTrackingCmds(trackerID, logger)

	logger.Info("Set tracker status to learning finished")
	err = GetTrackerManager().SetTrackerStatus(trackerID, models.StatusLearningFinished)
	if err != nil {
		return
	}
}

func (mgr *LearningManager) learningCreateTrackingCmds(trackerID int, logger *log.Entry) {
	logger.Info("Start creating tracking commands")

	trackCmd := &models.Command{
		TrackerID:    trackerID,
		Command:      models.CmdSendTrackingInformation,
		SleepTimeSec: mgr.sleepBetweenLearnSec,
	}

	for it := 0; it < mgr.learnCount; it++ {
		trackCmd.ID = 0
		GetTrackerManager().AddTrackerCommand(trackCmd)

		time.Sleep(time.Duration(mgr.sleepBetweenLearnSec-1) * time.Second)
	}
	logger.Info("Finished creating tracking commands")
}

func (mgr *LearningManager) NewTrackingData(trackerID int, scanRes []*models.ScanResult) (err error) {
	trackingDataLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		trackingDataLog.WithField("err", err).Error("Failed to get tracker with tracker ID")
		return
	}

	switch tracker.Status {
	case models.StatusIdle, models.StatusLearningFinished:
		err = errors.New("No tracking data expected")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unexpected tracking data")
	case models.StatusLearning:
		err = mgr.newLearningTrackingData(trackerID, scanRes)
	case models.StatusTracking:
		err = errors.New("Not implemented yes") //TODO
	}

	if err != nil {
		log.Error(err)
	}

	return
}

func (mgr *LearningManager) newLearningTrackingData(trackerID int, scanRes []*models.ScanResult) (err error) {
	for _, scan := range scanRes {
		scan.TrackerID = trackerID
	}

	err = mgr.scanResultRep.CreateAll(scanRes)
	return
}

func (mgr *LearningManager) FinishLearning(trackerID, roomID int, ssids []string) (err error) {
	finishLearningLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to get tracker")
		err = fmt.Errorf("tracker: %v", err)
		return
	}

	if tracker.Status != models.StatusLearningFinished {
		err = errors.New("Tracker is not in status LearningFinished")
		return
	}

	room, err := GetRoomManager().GetRoomByID(roomID)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to get room")
		err = fmt.Errorf("room: %v", err)
		return
	}

	//TODO: Calc something useful from saved scans

	err = GetRoomManager().SetRoomLearned(room.ID, true)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to save room")
		err = fmt.Errorf("room: %v", err)
		return
	}

	err = GetTrackerManager().SetTrackerStatus(tracker.ID, models.StatusIdle)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to set tracker status to idle after learning")
		return
	}

	return
}

func (mgr *LearningManager) GetLearningStatus(trackerID int) (done bool, ssids []string, err error) {
	learningStatusLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		learningStatusLog.WithField("err", err).Error("Failed to get tracker")
		return
	} else if tracker.Status != models.StatusLearning && tracker.Status != models.StatusLearningFinished {
		err = errors.New("Tracker currently not in learning or learning finished status")
		return
	}

	done = tracker.Status == models.StatusLearningFinished

	scanRes, err := mgr.scanResultRep.GetAllForTracker(trackerID)
	if err != nil && !mgr.scanResultRep.IsRecordNotFoundError(err) {
		learningStatusLog.WithField("err", err).Error("Failed to get scan results for tracker")
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
