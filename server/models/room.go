package models

type RoomID int

type Room struct {
	ID           RoomID              `json:"id"`
	Label        string              `json:"label"`
	IsLearned    bool                `json:"is_learned"`
	TrackingData []BSSIDTrackingData `json:"tracking_data"`
}
