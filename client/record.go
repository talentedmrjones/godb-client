package client

import (
	"github.com/twinj/uuid"
)

// Create writes a new record to the server. If id field was not set it will set a UUID. It returns a Reply
func (record Record) Create () *Reply {

	// ensure data includes "id" key
	if _, dataIdExists := record.Data["id"]; !dataIdExists {
		// if not make it unique
		record.SetString("id", uuid.NewV4().String())
	}

	record.action = "c"

	return transmit(record.Patient)
}

// Update writes an existing record to the server and returns a Reply
func (record Record) Update () *Reply {
	record.action = "u"
	return transmit(record.Patient)
}
