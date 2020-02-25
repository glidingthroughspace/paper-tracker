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

func (rep *GormRoomRepository) GetByID(roomID models.RoomID) (room *models.Room, err error) {
	room = &models.Room{}
	err = databaseConnection.First(room, &models.Room{ID: roomID}).Error
	return
}

func (rep *GormRoomRepository) GetAll() (rooms []*models.Room, err error) {
	err = databaseConnection.Find(&rooms).Error
	return
}

func (rep *GormRoomRepository) Delete(roomID models.RoomID) (err error) {
	err = databaseConnection.Delete(&models.Room{ID: roomID}).Error
	return
}

func (rep *GormRoomRepository) Update(room *models.Room) (err error) {
	err = databaseConnection.Save(room).Error
	return
}

func (rep *GormRoomRepository) SetLearnedByID(roomID models.RoomID, learned bool) (err error) {
	err = databaseConnection.Model(&models.Room{}).Where(&models.Room{ID: roomID}).Update("is_learned", learned).Error
	return
}
