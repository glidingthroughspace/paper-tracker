package models

type TrackerResponse struct {
	BatteryPercentage float32
}

type TrackingInformationResponse struct {
	TrackerResponse
	ScanResults []ScanResult
}
