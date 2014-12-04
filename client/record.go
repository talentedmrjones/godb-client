package client

import (
	//"errors"
	"encoding/binary"
	"github.com/twinj/uuid"
)

func (record *Record) String(field string) string {
	return string(record.Data[field])
}

func (record *Record) SetString (field string, value string) {
	record.Data[field] = []byte(value)
}

func (record *Record) SetUint32 (field string, value uint32) {
	valueBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(valueBuffer, value)
	record.Data[field] = valueBuffer
}

// Create write a record to the server
func (record *Record) Create () *Reply {

	// ensure data includes "id" key
	if _, dataIdExists := record.Data["id"]; !dataIdExists {
		// if not make it unique
		record.SetString("id", uuid.NewV4().String())
	}

	record.action = "c"

	reply := transmit(record)
	return reply

}


// Read querys the server
func (record *Record) Read () *Reply {

	record.action = "r"

	reply := transmit(record)

	return reply

}
