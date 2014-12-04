package client

import (
	//"errors"
	"github.com/twinj/uuid"
)



// Create write a record to the server
func (record *Record) Create () *Reply {

	// ensure data includes "id" key
	if _, dataIdExists := record.Data["id"]; !dataIdExists {
		// if not make it unique
		record.SetString("id", uuid.NewV4().String())
	}

	record.action = "c"

	reply := transmit(&record.Patient)
	return reply

}
