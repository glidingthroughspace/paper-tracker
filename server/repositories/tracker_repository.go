package repositories

import "paper-tracker/models"

type TrackerRepository interface {
	Create(tracker *models.Tracker) error
	GetAll() ([]*models.Tracker, error)
	GetByID(trackerID int) (*models.Tracker, error)
	Update(tracker *models.Tracker) error
	IsRecordNotFoundError(err error) bool
}
