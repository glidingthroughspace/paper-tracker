package gorm

import "paper-tracker/models"

func init() {
	databaseModels = append(databaseModels, &models.BSSIDTrackingData{})
}
