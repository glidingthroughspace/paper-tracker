package communication

import "paper-tracker/models"

type TrackerCmdResponse struct {
	BatteryPercentage int  `json:"battery_percentage"`
	IsCharging        bool `json:"is_charging"`
}

type TrackingCmdResponse struct {
	TrackerCmdResponse `json:"tracker_cmd_response"`
	IsLastBatch        bool                 `json:"is_last_batch"`
	ScanResults        []*models.ScanResult `json:"scan_results"`
}
