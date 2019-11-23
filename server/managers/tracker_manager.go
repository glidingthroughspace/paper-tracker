package managers

import (
	"errors"
	"fmt"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

	log "github.com/sirupsen/logrus"
)

var defaultSleepCmd *models.Command

type TrackerManager struct {
	trackerRep           repositories.TrackerRepository
	cmdRep               repositories.CommandRepository
	scanResultRep        repositories.ScanResultRepository
	roomRep              repositories.RoomRepository
	learnCount           int
	sleepBetweenLearnSec int
	done                 chan struct{}
}

func CreateTrackerManager(trackerRep repositories.TrackerRepository, cmdRep repositories.CommandRepository, scanResultRep repositories.ScanResultRepository, roomRep repositories.RoomRepository, defaultSleepSec, learnCount, sleepBetweenLearnSec int) *TrackerManager {
	trackerManager := &TrackerManager{
		trackerRep:           trackerRep,
		cmdRep:               cmdRep,
		scanResultRep:        scanResultRep,
		roomRep:              roomRep,
		learnCount:           learnCount,
		sleepBetweenLearnSec: sleepBetweenLearnSec,
		done:                 make(chan struct{}),
	}

	defaultSleepCmd = &models.Command{
		Command:      models.CmdSleep,
		SleepTimeSec: defaultSleepSec,
	}

	return trackerManager
}

func (mgr *TrackerManager) GetAllTrackers() (trackers []*models.Tracker, err error) {
	trackers, err = mgr.trackerRep.GetAll()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all trackers")
		return
	}
	return
}

func (mgr *TrackerManager) NotifyNewTracker() (tracker *models.Tracker, err error) {
	tracker = &models.Tracker{Label: "New Tracker"}
	err = mgr.trackerRep.Create(tracker)
	if err != nil {
		log.WithField("err", err).Error("Failed to create new tracker")
		return
	}
	return
}

func (mgr *TrackerManager) PollCommand(trackerID int) (cmd *models.Command, err error) {
	pollLog := log.WithField("trackerID", trackerID)

	_, err = mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		pollLog.WithField("err", err).Error("Failed to get tracker with tracker ID")
		return
	}

	cmd, err = mgr.cmdRep.GetNextCommand(trackerID)
	if err != nil && !mgr.cmdRep.IsRecordNotFoundError(err) {
		pollLog.WithField("err", err).Error("Failed to get next command for tracker")
		return
	} else if mgr.cmdRep.IsRecordNotFoundError(err) {
		pollLog.Info("No command for tracker, return default sleep")
		err = nil
		cmd = defaultSleepCmd
		return
	}

	err = mgr.cmdRep.Delete(cmd.ID)
	if err != nil {
		pollLog.WithField("err", err).Error("Failed to delete command")
		return
	}

	if _, nextErr := mgr.cmdRep.GetNextCommand(trackerID); !mgr.cmdRep.IsRecordNotFoundError(nextErr) {
		cmd.SleepTimeSec = 0
	}

	return
}

//TODO: Move learning to own manager
func (mgr *TrackerManager) StartLearning(trackerID int) (learnTimeSec int, err error) {
	learnLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		learnLog.WithField("err", err).Error("Failed to get tracker with tracker ID")
		return
	} else if tracker.Status != models.StatusIdle {
		err = errors.New("Given tracker is not in idle mode")
		return
	}

	go mgr.learningRoutine(tracker, learnLog)

	learnTimeSec = mgr.learnCount * mgr.sleepBetweenLearnSec
	return
}

func (mgr *TrackerManager) learningRoutine(tracker *models.Tracker, logger *log.Entry) {
	logger.Trace("Start routine")

	logger.Trace("Set tracker status to learning")
	tracker.Status = models.StatusLearning
	err := mgr.trackerRep.Update(tracker)
	if err != nil {
		logger.WithField("err", err).Error("Failed to update tracker with status learning")
		return
	}

	mgr.learningCreateTrackingCmds(tracker, logger)

	logger.Trace("Set tracker status to learning finished")
	tracker.Status = models.StatusLearningFinished
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		logger.WithField("err", err).Error("Failed to update tracker with status learning finished")
		return
	}
}

func (mgr *TrackerManager) learningCreateTrackingCmds(tracker *models.Tracker, logger *log.Entry) {
	logger.Info("Start creating tracking commands")

	trackCmd := &models.Command{
		TrackerID:    tracker.ID,
		Command:      models.CmdSendTrackingInformation,
		SleepTimeSec: mgr.sleepBetweenLearnSec,
	}

	for it := 0; it < mgr.learnCount; it++ {
		err := mgr.cmdRep.Create(trackCmd)
		if err != nil {
			logger.WithField("err", err).Error("Failed to insert tracking command")
		}

		time.Sleep(time.Duration(mgr.sleepBetweenLearnSec-1) * time.Second)
	}
	logger.Info("Finished creating tracking commands, set tracker status to idle")
}

func (mgr *TrackerManager) NewTrackingData(trackerID int, scanRes []*models.ScanResult) (err error) {
	trackingDataLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.trackerRep.GetByID(trackerID)
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
	return
}

func (mgr *TrackerManager) newLearningTrackingData(trackerID int, scanRes []*models.ScanResult) (err error) {
	for _, scan := range scanRes {
		scan.TrackerID = trackerID
	}

	err = mgr.scanResultRep.CreateAll(scanRes)
	return
}

func (mgr *TrackerManager) FinishLearning(trackerID, roomID int) (err error) {
	finishLearningLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to get tracker")
		err = fmt.Errorf("tracker: %v", err)
		return
	}

	if tracker.Status != models.StatusLearningFinished {
		err = errors.New("Tracker is not in status LearningFinished")
		return
	}

	_, err = mgr.roomRep.GetByID(roomID)
	if err != nil {
		finishLearningLog.WithField("err", err).Error("Failed to get room")
		err = fmt.Errorf("room: %v", err)
		return
	}

	//TODO: Calc something useful from saved scans

	return
}
