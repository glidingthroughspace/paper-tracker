package communication

import "paper-tracker/models"

type TrackerCmdResponse struct {
	BatteryPercentage float32 `json:"battery_percentage"`
}

type TrackingCmdResponse struct {
	TrackerCmdResponse `json:"tracker_cmd_response"`
	ScanResults        []*models.ScanResult `json:"scan_results"`
}
