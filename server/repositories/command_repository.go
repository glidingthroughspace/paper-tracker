package repositories

import "paper-tracker/models"

// Add used models to enable auto migration for them
func init() {
	databaseModels = append(databaseModels, &models.Command{})
}

type CommandRepository struct{}

func CreateCommandRepository() (*CommandRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &CommandRepository{}, nil
}

func (rep *CommandRepository) Create(command *models.Command) (err error) {
	err = databaseConnection.Create(command).Error
	return
}

func (rep *CommandRepository) GetNextCommand(trackerID int) (cmd *models.Command, err error) {
	cmd = &models.Command{}
	err = databaseConnection.Where(&models.Command{TrackerID: trackerID}).Order("created_at asc").First(cmd).Error
	return
}

func (rep *CommandRepository) Delete(trackerID int) (err error) {
	err = databaseConnection.Delete(&models.Command{ID: trackerID}).Error
	return
}
