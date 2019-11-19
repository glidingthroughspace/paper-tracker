package repositories

import "paper-tracker/models"

type CommandRepository interface {
	Create(command *models.Command) (err error)
	GetNextCommand(trackerID int) (cmd *models.Command, err error)
	Delete(trackerID int) (err error)
}
