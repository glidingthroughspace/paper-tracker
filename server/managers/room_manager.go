package managers

import (
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var roomManager *RoomManager

type RoomManager struct {
	roomRep repositories.RoomRepository
}

func CreateRoomManager(roomRep repositories.RoomRepository) *RoomManager {
	if roomManager != nil {
		return roomManager
	}

	roomManager = &RoomManager{
		roomRep: roomRep,
	}

	return roomManager
}

func GetRoomManager() *RoomManager {
	return roomManager
}

func (mgr *RoomManager) CreateRoom(room *models.Room) (err error) {
	room.ID = 0
	err = mgr.roomRep.Create(room)
	if err != nil {
		log.WithFields(log.Fields{"room": room, "err": err}).Error("Failed to create room")
		return
	}
	return
}

func (mgr *RoomManager) GetRoomByID(roomID int) (room *models.Room, err error) {
	room, err = mgr.roomRep.GetByID(roomID)
	if err != nil {
		log.WithFields(log.Fields{"roomID": roomID, "err": err}).Error("Failed to get room")
		return
	}
	return
}

func (mgr *RoomManager) GetAllRooms() (rooms []*models.Room, err error) {
	rooms, err = mgr.roomRep.GetAll()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all rooms")
		return
	}
	return
}

func (mgr *RoomManager) SetRoomLearned(roomID int, learned bool) (err error) {
	err = mgr.roomRep.SetLearnedByID(roomID, learned)
	if err != nil {
		log.WithFields(log.Fields{"roomID": roomID, "learned": learned, "err": err}).Error("Failed to set room learned")
		return
	}
	return
}

func (mgr *RoomManager) UpdateRoom(room *models.Room) (err error) {
	err = mgr.roomRep.Update(room)
	if err != nil {
		log.WithFields(log.Fields{"roomID": room.ID, "err": err}).Error("Failed to update room")
		return
	}
	return
}
