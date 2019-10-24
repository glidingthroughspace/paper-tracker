package managers

import (
	"fmt"
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

type TrackerManager struct {
	trackerRep *repositories.TrackerRepository
	done       chan struct{}
}

var trackerManager *TrackerManager

func CreateTrackerManager(trackerRep *repositories.TrackerRepository) *TrackerManager {
	if trackerManager != nil {
		return trackerManager
	}

	trackerManager = &TrackerManager{
		trackerRep: trackerRep,
		done:       make(chan struct{}),
	}
	return trackerManager
}

func GetTrackerManager() *TrackerManager {
	return trackerManager
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
