package managers

import (
	"errors"
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var defaultSleepCmd *models.Command

type TrackerManager struct {
	trackerRep repositories.TrackerRepository
	cmdRep     repositories.CommandRepository
	done       chan struct{}
}

func CreateTrackerManager(trackerRep repositories.TrackerRepository, cmdRep repositories.CommandRepository, defaultSleepSec int) *TrackerManager {
	trackerManager := &TrackerManager{
		trackerRep: trackerRep,
		cmdRep:     cmdRep,
		done:       make(chan struct{}),
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
	if err != nil && !repositories.IsRecordNotFoundError(err) {
		pollLog.WithField("err", err).Error("Failed to get next command for tracker")
		return
	} else if repositories.IsRecordNotFoundError(err) {
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

	if _, err = mgr.cmdRep.GetNextCommand(trackerID); !repositories.IsRecordNotFoundError(err) {
		cmd.SleepTimeSec = 0
	} else {
		err = nil
	}

	return
}

func (mgr *TrackerManager) StartLearning(trackerID int) (err error) {
	learnLog := log.WithField("trackerID", trackerID)

	tracker, err := mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		learnLog.WithField("err", err).Error("Failed to get tracker with tracker ID")
		return
	} else if tracker.Status != models.StatusIdle {
		err = errors.New("Given tracker is not in idle mode")
		return
	}

	tracker.Status = models.StatusLearning
	err = mgr.trackerRep.Update(tracker)
	if err != nil {
		learnLog.WithField("err", err).Error("Failed to update tracker with status learning")
		return
	}

	return
}
