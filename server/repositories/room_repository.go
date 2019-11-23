package repositories

import "paper-tracker/models"

type RoomRepository interface {
	Create(room *models.Room) error
	GetAll() ([]*models.Room, error)
	Delete(roomID int) error
	IsRecordNotFoundError(err error) bool
}
