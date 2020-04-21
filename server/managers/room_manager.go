package managers

import (
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var roomManager RoomManager

type RoomManager interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(roomID models.RoomID) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error)
	SetRoomLearned(roomID models.RoomID, learned bool) error
	UpdateRoom(room *models.Room) error
	DeleteRoom(roomID models.RoomID) error
}

type RoomManagerImpl struct {
	roomRep repositories.RoomRepository
}

func CreateRoomManager(roomRep repositories.RoomRepository) RoomManager {
	if roomManager != nil {
		return roomManager
	}

	roomManager = &RoomManagerImpl{
		roomRep: roomRep,
	}

	return roomManager
}

func GetRoomManager() RoomManager {
	return roomManager
}

func (mgr *RoomManagerImpl) CreateRoom(room *models.Room) (err error) {
	room.ID = 0
	err = mgr.roomRep.Create(room)
	if err != nil {
		log.WithFields(log.Fields{"room": room, "err": err}).Error("Failed to create room")
		return
	}
	return
}

func (mgr *RoomManagerImpl) GetRoomByID(roomID models.RoomID) (room *models.Room, err error) {
	room, err = mgr.roomRep.GetByID(roomID)
	if err != nil {
		log.WithFields(log.Fields{"roomID": roomID, "err": err}).Error("Failed to get room")
		return
	}

	err = mgr.fillRoomInfo(room)
	return
}

func (mgr *RoomManagerImpl) GetAllRooms() (rooms []*models.Room, err error) {
	rooms, err = mgr.roomRep.GetAll()
	if err != nil {
		log.WithError(err).Error("Failed to get all rooms")
		return
	}

	for _, room := range rooms {
		err = mgr.fillRoomInfo(room)
		if err != nil {
			log.WithFields(log.Fields{"err": err, "roomID": room.ID}).Error("Failed to fill room infos for list")
			continue
		}
	}

	return
}

func (mgr *RoomManagerImpl) fillRoomInfo(room *models.Room) (err error) {
	fillInfoLog := log.WithField("roomID", room.ID)

	stepCount, err := GetWorkflowTemplateManager().NumberOfStepsReferringToRoom(room.ID)
	if err != nil {
		fillInfoLog.WithError(err).Error("Failed to get step count for room")
		return
	}
	if stepCount > 0 {
		room.DeleteLocked = true
	} else {
		room.DeleteLocked = false
	}

	return
}

func (mgr *RoomManagerImpl) SetRoomLearned(roomID models.RoomID, learned bool) (err error) {
	err = mgr.roomRep.SetLearnedByID(roomID, learned)
	if err != nil {
		log.WithFields(log.Fields{"roomID": roomID, "learned": learned, "err": err}).Error("Failed to set room learned")
		return
	}
	return
}

func (mgr *RoomManagerImpl) UpdateRoom(room *models.Room) (err error) {
	err = mgr.roomRep.Update(room)
	if err != nil {
		log.WithFields(log.Fields{"roomID": room.ID, "err": err}).Error("Failed to update room")
		return
	}
	return
}

func (mgr *RoomManagerImpl) DeleteRoom(roomID models.RoomID) (err error) {
	deleteLog := log.WithField("roomID", roomID)

	err = mgr.roomRep.Delete(roomID)
	if err != nil {
		deleteLog.WithError(err).Error("Failed to delete room")
		return
	}
	return
}
