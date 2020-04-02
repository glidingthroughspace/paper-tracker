package communication

import "paper-tracker/models"

type TrackerCmdResponse struct {
	BatteryPercentage int  `codec:"battery_percentage"`
	IsCharging        bool `codec:"is_charging"`
}

type TrackingCmdResponse struct {
	TrackerCmdResponse
	// ResultID is an ID for identifying batches comprising a single result
	ResultID uint64 `codec:"result_id"`
	// BatchCount batches are in the given result
	BatchCount  uint8                `codec:"result_batch_count"`
	ScanResults []*models.ScanResult `codec:"scan_results"`
}
