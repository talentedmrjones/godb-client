package client

import (
	"github.com/twinj/uuid"
)

func NewCommand (command string, patient *Patient) *Command {

	id := uuid.NewV4()
	return &Command{
		id.String(),
		command,
		patient.table.db.name,
		patient.table.name,
		patient.Data,
	}
}
