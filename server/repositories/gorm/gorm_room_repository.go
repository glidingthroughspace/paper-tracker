package gorm

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.Room{})
}

type GormRoomRepository struct{}

func CreateGormRoomRepository() (*GormRoomRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormRoomRepository{}, nil
}

func (rep *GormRoomRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormRoomRepository) Create(room *models.Room) (err error) {
	err = databaseConnection.Create(room).Error
	return
}

func (rep *GormRoomRepository) GetByID(roomID int) (room *models.Room, err error) {
	room = &models.Room{}
	err = databaseConnection.First(room, &models.Room{ID: roomID}).Error
	return
}

func (rep *GormRoomRepository) GetAll() (rooms []*models.Room, err error) {
	err = databaseConnection.Find(&rooms).Error
	return
}

func (rep *GormRoomRepository) Delete(roomID int) (err error) {
	err = databaseConnection.Delete(&models.Room{ID: roomID}).Error
	return
}