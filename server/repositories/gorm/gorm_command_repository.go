package gorm

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.Command{})
}

type GormCommandRepository struct{}

func CreateGormCommandRepository() (*GormCommandRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormCommandRepository{}, nil
}

func (rep *GormCommandRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormCommandRepository) Create(command *models.Command) (err error) {
	err = databaseConnection.Create(command).Error
	return
}

func (rep *GormCommandRepository) GetNextCommand(trackerID int) (cmd *models.Command, err error) {
	cmd = &models.Command{}
	err = databaseConnection.Where(&models.Command{TrackerID: trackerID}).Order("created_at asc").First(cmd).Error
	return
}

func (rep *GormCommandRepository) Delete(trackerID int) (err error) {
	err = databaseConnection.Delete(&models.Command{ID: trackerID}).Error
	return
}
