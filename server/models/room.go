package models

type Room struct {
	ID           int                 `json:"id,omitempty"`
	Label        string              `json:"label,omitempty"`
	IsLearned    bool                `json:"is_learned"`
	TrackingData []BSSIDTrackingData `json:"tracking_data,omitempty"`
}
