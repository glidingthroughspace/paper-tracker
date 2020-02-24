package repositories

import "paper-tracker/models"

type CommandRepository interface {
	Create(command *models.Command) error
	GetNextCommand(trackerID models.TrackerID) (*models.Command, error)
	Delete(commandID models.CommandID) error
	IsRecordNotFoundError(err error) bool
}
