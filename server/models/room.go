package models

type RoomID int

type Room struct {
	ID           RoomID              `json:"id" gorm:"primary_key;auto_increment"`
	Label        string              `json:"label"`
	IsLearned    bool                `json:"is_learned"`
	TrackingData []BSSIDTrackingData `json:"tracking_data"`
}

type ScoredRoom struct {
	Room
	Score float64 `json:"score"`
}

func ScoredRoomFromRoom(room *Room, score float64) *ScoredRoom {
	return &ScoredRoom{
		Room:  *room,
		Score: score,
	}
}
