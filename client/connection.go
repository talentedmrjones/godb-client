package client

import (
	"net"
)

type Connection struct {
	conn net.Conn
}


func NewConnection (address, port string) (*Connection) {

	conn, err := net.Dial("tcp", address+":"+port)
	if err!=nil {
		// handle error here
	}

	return &Connection{conn}
}

func (con *Connection) Database (dbName string) *Database {
	return &Database{con, dbName}
}
