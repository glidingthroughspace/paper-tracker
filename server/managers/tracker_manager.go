package managers

import (
	"fmt"
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var defaultSleepCmd *models.Command

type TrackerManager struct {
	trackerRep *repositories.TrackerRepository
	cmdRep     *repositories.CommandRepository
	done       chan struct{}
}

var trackerManager *TrackerManager

func CreateTrackerManager(trackerRep *repositories.TrackerRepository, cmdRep *repositories.CommandRepository, defaultSleepSec int) *TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManager{
		trackerRep: trackerRep,
		done:       make(chan struct{}),
	}

	defaultSleepCmd = &models.Command{
		Command:      models.Sleep,
		SleepTimeSec: defaultSleepSec,
	}

	return trackerManager
}

func GetTrackerManager() *TrackerManager {
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
		return nil, fmt.Errorf("Failed to create new tracker")
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
	}

	return
}
