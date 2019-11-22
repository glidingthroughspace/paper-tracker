package repositories

import "paper-tracker/models"

type TrackerRepository interface {
	Create(tracker *models.Tracker) (err error)
	GetAll() (trackers []*models.Tracker, err error)
	GetByID(trackerID int) (tracker *models.Tracker, err error)
	Update(tracker *models.Tracker) (err error)
}
