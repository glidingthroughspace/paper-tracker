package managers

import (
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var defaultSleepCmd *models.Command
var trackerManager *TrackerManager

type TrackerManager struct {
	trackerRep repositories.TrackerRepository
	cmdRep     repositories.CommandRepository
	done       chan struct{}
}

func CreateTrackerManager(trackerRep repositories.TrackerRepository, cmdRep repositories.CommandRepository, defaultSleepSec int) *TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManager{
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

func GetTrackerManager() *TrackerManager {
	return trackerManager
}

func (mgr *TrackerManager) GetTrackerByID(trackerID int) (tracker *models.Tracker, err error) {
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
		log.WithField("err", err).Error("Failed to get all trackers")
		return
	}
	return
}

func (mgr *TrackerManager) SetTrackerStatus(trackerID int, status models.TrackerStatus) (err error) {
	err = mgr.trackerRep.SetStatusByID(trackerID, status)
	if err != nil {
		log.WithFields(log.Fields{"trackerID": trackerID, "status": status, "err": err}).Error("Failed to set status of tracker")
		return
	}
	return
}

func (mgr *TrackerManager) AddTrackerCommand(command *models.Command) (err error) {
	err = mgr.cmdRep.Create(command)
	if err != nil {
		log.WithFields(log.Fields{"command": command, "err": err}).Error("Failed to add tracker command")
		return
	}
	return
}

func (mgr *TrackerManager) NotifyNewTracker() (tracker *models.Tracker, err error) {
	tracker = &models.Tracker{Label: "New Tracker", Status: models.StatusIdle}
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
