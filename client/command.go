package client

import (
	"github.com/twinj/uuid"
)

func NewCommand (command string, record *Record) *Command {

	id := uuid.NewV4()
	return &Command{
		id.String(),
		command,
		record.table.db.name,
		record.table.name,
		record.Data,
	}
}
