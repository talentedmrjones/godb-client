package client

import (
	"errors"
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

func (record *Record) Create () error {
	// ensure data includes "id" key

	if _, dataIdExists := record.data["id"]; !dataIdExists {
		err := errors.New("ID_MISSING")
		return err
	}

	command := NewCommand("c", record)
	transmit(record.table.db.connection.conn, command)

	return nil
}
