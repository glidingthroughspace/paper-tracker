package managers

import (
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

type RoomManager struct {
	roomRep repositories.RoomRepository
}

func CreateRoomManager(roomRep repositories.RoomRepository) *RoomManager {
	roomManager := &RoomManager{
		roomRep: roomRep,
	}

	return roomManager
}

func (mgr *RoomManager) GetAllRooms() (rooms []*models.Room, err error) {
	rooms, err = mgr.roomRep.GetAll()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all rooms")
		return
	}
	return
}
