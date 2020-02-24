package gorm

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.ScanResult{})
}

type GormScanResultRepository struct{}

func CreateGormScanResultRepository() (*GormScanResultRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormScanResultRepository{}, nil
}

func (rep *GormScanResultRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormScanResultRepository) CreateAll(scanRes []*models.ScanResult) (err error) {
	for _, scan := range scanRes {
		err = databaseConnection.Create(scan).Error
		if err != nil {
			return
		}
	}
	return
}

func (rep *GormScanResultRepository) GetAllForTracker(trackerID models.TrackerID) (scanRes []*models.ScanResult, err error) {
	err = databaseConnection.Find(&scanRes, &models.ScanResult{TrackerID: trackerID}).Error
	return
}

func (rep *GormScanResultRepository) DeleteForTracker(trackerID models.TrackerID) (err error) {
	err = databaseConnection.Where(&models.ScanResult{TrackerID: trackerID}).Delete(&models.ScanResult{}).Error
	return
}
