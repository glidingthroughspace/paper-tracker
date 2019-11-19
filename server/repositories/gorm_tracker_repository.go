package repositories

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.Tracker{})
}

type GormTrackerRepository struct{}

func CreateGormTrackerRepository() (*GormTrackerRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormTrackerRepository{}, nil
}

func (rep *GormTrackerRepository) Create(tracker *models.Tracker) (err error) {
	err = databaseConnection.Create(tracker).Error
	return
}

func (rep *GormTrackerRepository) GetAll() (trackers []*models.Tracker, err error) {
	err = databaseConnection.Find(&trackers).Error
	return
}

func (rep *GormTrackerRepository) GetByID(trackerID int) (tracker *models.Tracker, err error) {
	tracker = &models.Tracker{}
	err = databaseConnection.First(tracker, &models.Tracker{ID: trackerID}).Error
	return
}
