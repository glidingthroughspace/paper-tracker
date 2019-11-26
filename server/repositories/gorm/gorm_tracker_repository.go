package gorm

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

func (rep *GormTrackerRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
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

func (rep *GormTrackerRepository) Update(tracker *models.Tracker) (err error) {
	err = databaseConnection.Save(tracker).Error
	return
}

func (rep *GormTrackerRepository) SetStatusByID(trackerID int, status models.TrackerStatus) (err error) {
	err = databaseConnection.Where(&models.Tracker{ID: trackerID}).Updates(&models.Tracker{Status: status}).Error
	return
}
