package repositories

import "paper-tracker/models"

type CommandRepository interface {
	Create(command *models.Command) error
	GetNextCommand(trackerID int) (*models.Command, error)
	Delete(trackerID int) error
	IsRecordNotFoundError(err error) bool
}
