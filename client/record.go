package client

import (
	"github.com/twinj/uuid"
)

// Record ...
type Record struct {
	table  *Table
	Data   map[string]interface{}
	action string `default:""`
}

// Records ...
type Records []Record

// SetField ...
func (record Record) SetField(field string, value interface{}) {
	record.Data[field] = value
}

// Create writes a new record to the server. If id field was not set it will set a UUID. It returns a Reply
func (record Record) Create() {

	// ensure data includes "id" key
	if _, dataIDExists := record.Data["id"]; !dataIDExists {
		// if not make it unique
		record.SetField("id", uuid.NewV4().String())
	}

	record.action = "c"

	record.table.db.connection.records <- record
}
