package client

import (
	//"errors"
	"encoding/binary"
	"github.com/twinj/uuid"
)

type Record struct {
	table		*Table
	data		map[string][]byte
	action	string
}

func (record *Record) SetFieldString (field string, value string) {
	record.data[field] = []byte(value)
}

func (record *Record) SetFieldUint32 (field string, value uint32) {
	valueBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(valueBuffer, value)
	record.data[field] = valueBuffer
}

func (record *Record) Create () (*Reply, *Record) {

	// ensure data includes "id" key
	if _, dataIdExists := record.data["id"]; !dataIdExists {
		// if not make it unique
		record.SetFieldString("id", uuid.NewV4().String())
	}

	record.action = "c"

	reply := transmit(record)

	return reply, record

}
