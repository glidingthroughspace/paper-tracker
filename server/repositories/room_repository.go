package repositories

import "paper-tracker/models"

type RoomRepository interface {
	Create(room *models.Room) error
	GetByID(roomID models.RoomID) (*models.Room, error)
	GetAll() ([]*models.Room, error)
	Delete(roomID models.RoomID) error
	SetLearnedByID(roomID models.RoomID, learned bool) error
	Update(room *models.Room) error
	IsRecordNotFoundError(err error) bool
}
