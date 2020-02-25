package repositories

import "paper-tracker/models"

type ScanResultRepository interface {
	CreateAll(scanRes []*models.ScanResult) error
	GetAllForTracker(trackerID models.TrackerID) ([]*models.ScanResult, error)
	DeleteForTracker(trackerID models.TrackerID) error
	IsRecordNotFoundError(err error) bool
}
