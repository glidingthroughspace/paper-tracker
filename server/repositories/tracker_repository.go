package repositories

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.Tracker{})
}

type TrackerRepository struct{}

func CreateTrackerRepository() (*TrackerRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &TrackerRepository{}, nil
}

func (rep *TrackerRepository) Create(tracker *models.Tracker) (err error) {
	err = databaseConnection.Create(tracker).Error
	return
}
