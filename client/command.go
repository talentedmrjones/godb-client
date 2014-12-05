package client

import (
	"github.com/twinj/uuid"
)

func NewCommand (command string, patient Patient) Command {

	return Command{
		uuid.NewV4().String(),
		command,
		patient.table.db.name,
		patient.table.name,
		patient.Data,
	}
}
