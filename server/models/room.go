package models

type Room struct {
	ID           int                 `json:"id"`
	Label        string              `json:"label"`
	IsLearned    bool                `json:"is_learned"`
	TrackingData []BSSIDTrackingData `json:"tracking_data"`
}
