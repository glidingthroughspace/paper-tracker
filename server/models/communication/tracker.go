package communication

import "paper-tracker/models"

type TrackerCmdResponse struct {
	BatteryPercentage float32
}

type TrackingCmdResponse struct {
	TrackerCmdResponse
	ScanResults []*models.ScanResult
}
