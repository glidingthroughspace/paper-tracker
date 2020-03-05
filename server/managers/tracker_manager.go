package managers

import (
	"errors"
	"paper-tracker/models"
	"paper-tracker/models/communication"
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

func (mgr *TrackerManager) AddTrackerCommand(command *models.Command) (err error) {
	err = mgr.cmdRep.Create(command)
	if err != nil {
		log.WithFields(log.Fields{"command": command, "err": err}).Error("Failed to add tracker command")
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

	_, err = mgr.trackerRep.GetByID(trackerID)
	if err != nil {
		pollLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	cmd, err = mgr.cmdRep.GetNextCommand(trackerID)
	if err != nil && !mgr.cmdRep.IsRecordNotFoundError(err) {
		pollLog.WithError(err).Error("Failed to get next command for tracker")
		return
	} else if mgr.cmdRep.IsRecordNotFoundError(err) {
		pollLog.Info("No command for tracker, return default sleep")
		err = nil
		cmd = defaultSleepCmd
		return
	}

	err = mgr.cmdRep.Delete(cmd.ID)
	if err != nil {
		pollLog.WithError(err).Error("Failed to delete command")
		return
	}

	if _, nextErr := mgr.cmdRep.GetNextCommand(trackerID); !mgr.cmdRep.IsRecordNotFoundError(nextErr) {
		cmd.SleepTimeSec = 0
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

func (mgr *TrackerManager) NewTrackingData(trackerID models.TrackerID, scanRes []*models.ScanResult) (err error) {
	trackingDataLog := log.WithField("trackerID", trackerID)

	tracker, err := GetTrackerManager().GetTrackerByID(trackerID)
	if err != nil {
		trackingDataLog.WithError(err).Error("Failed to get tracker with tracker ID")
		return
	}

	switch tracker.Status {
	case models.TrackerStatusIdle, models.TrackerStatusLearningFinished:
		err = errors.New("No tracking data expected")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unexpected tracking data")
	case models.TrackerStatusLearning:
		err = GetLearningManager().newLearningTrackingData(trackerID, scanRes)
	case models.TrackerStatusTracking:
		err = errors.New("Not implemented yes") //TODO
	default:
		err = errors.New("Unknown tracker status")
		trackingDataLog.WithField("trackerStatus", tracker.Status).Error("Unknown tracker status")
	}

	if err != nil {
		log.Error(err)
	}

	return
}
