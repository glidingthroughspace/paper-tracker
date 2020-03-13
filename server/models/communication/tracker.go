package communication

import "paper-tracker/models"

type TrackerCmdResponse struct {
	BatteryPercentage int  `json:"battery_percentage"`
	IsCharging        bool `json:"is_charging"`
}

type TrackingCmdResponse struct {
	TrackerCmdResponse `json:"tracker_cmd_response"`
	// ResultID is an ID for identifying batches comprising a single result
	ResultID uint64 `json:"result_id"`
	// BatchCount batches are in the given result
	BatchCount  uint8                `json:"result_batch_count"`
	ScanResults []*models.ScanResult `json:"scan_results"`
}
