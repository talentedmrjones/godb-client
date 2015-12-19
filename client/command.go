package client

import (
	"github.com/twinj/uuid"
)

/*
Command forms a protocol to communicate with godb server. It is used to relay commands and data from connection to table channels
*/
type Command struct {
	ID     string                 `json:"id"`
	Action string                 `json:"action"`
	Db     string                 `json:"db"`
	Table  string                 `json:"table"`
	Data   map[string]interface{} `json:"data"`
}

/*
NewCommand creates a Command with a unique v4 UUID
*/
func NewCommand(record Record) Command {

	return Command{
		uuid.NewV4().String(),
		record.action,
		record.table.db.name,
		record.table.name,
		record.Data,
	}
}
