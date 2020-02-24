package repositories

import "paper-tracker/models"

type TrackerRepository interface {
	Create(tracker *models.Tracker) error
	GetAll() ([]*models.Tracker, error)
	GetByID(trackerID models.TrackerID) (*models.Tracker, error)
	Update(tracker *models.Tracker) error
	SetStatusByID(trackerID models.TrackerID, status models.TrackerStatus) error
	IsRecordNotFoundError(err error) bool
}
