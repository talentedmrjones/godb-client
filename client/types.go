package client

import (
	"net"
)

type Connection struct {
	socket net.Conn
	replies map[string]chan *Reply
}

type Command struct {
	Id			string
	Action 	string
	Db			string
	Table 	string
	Query		map[string][]byte
}

type Database struct {
	connection 	*Connection // TODO change to chan *Connection to support pooling
	name 				string
}

type Patient struct {
	table		*Table
	Data		map[string][]byte
	action	string
}

type Query struct {
	Patient
}

type Record struct {
	Patient
}

type Records []Record

// Reply is received from server
type Reply struct {
	Id			string
	Status  uint16
	Result	Records
	Error		string
}

type Table struct {
	db 		*Database
	name	string
}
