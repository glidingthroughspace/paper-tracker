package models

type RoomID int

type Room struct {
	ID           RoomID              `json:"id" gorm:"primary_key;auto_increment"`
	Label        string              `json:"label"`
	IsLearned    bool                `json:"is_learned"`
	TrackingData []BSSIDTrackingData `json:"tracking_data"`
	DeleteLocked bool                `json:"delete_locked" gorm:"-"`
}
