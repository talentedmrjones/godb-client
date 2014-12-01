package client

import (
	"github.com/twinj/uuid"
)

type Command struct {
	Id			string
	Action 	string
	Db			string
	Table 	string
	Data		map[string][]byte
}

func NewCommand (command string, record *Record) *Command {

	id := uuid.NewV1()
	return &Command{
		id.String(),
		command,
		record.table.db.name,
		record.table.name,
		record.data,
	}
}
