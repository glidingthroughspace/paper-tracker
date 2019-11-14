package models

type Room struct {
	ID           int
	Label        string
	TrackingData []BSSIDTrackingData
}
