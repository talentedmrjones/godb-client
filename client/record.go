package client

import (
	//"errors"
	"encoding/binary"
)

type Record struct {
	table	*Table
	data	map[string][]byte
}

func (record *Record) SetFieldString (field string, value string) {
	record.data[field] = []byte(value)
}

func (record *Record) SetFieldUint32 (field string, value uint32) {
	valueBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(valueBuffer, value)
	record.data[field] = valueBuffer
}

func (record *Record) Create () *Reply {
	// ensure data includes "id" key

	if _, dataIdExists := record.data["id"]; !dataIdExists {
		reply := &Reply{}
		reply.Error = "ID_MISSING"
		return reply
	}


	command := NewCommand("c", record)
	record.table.db.connection.replies[command.Id] = make(chan *Reply)
	transmit(record.table.db.connection.socket, command)

	reply := <- record.table.db.connection.replies[command.Id]

	return reply
}
