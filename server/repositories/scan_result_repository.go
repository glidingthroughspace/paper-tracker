package repositories

import "paper-tracker/models"

type ScanResultRepository interface {
	CreateAll(scanRes []*models.ScanResult) error
	DeleteForTracker(trackerID int) error
	IsRecordNotFoundError(err error) bool
}
