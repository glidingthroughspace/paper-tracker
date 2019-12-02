package models

type Room struct {
	ID           int
	Label        string
	IsLearned    bool
	TrackingData []BSSIDTrackingData
}
