package client

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/binary"
	"fmt"
	"net"

)

type Connection struct {
	socket net.Conn
}


func NewConnection (address, port string) (*Connection) {

	conn, err := net.Dial("tcp", address+":"+port)
	if err!=nil {
		// handle error here
	}

	connection := &Connection{conn}

	return connection
}

func (connection *Connection) Database (dbName string) *Database {
	return &Database{connection, dbName}
}

// Receive continuously looks for data from the socket and relays that to a table's command channel
func (connection *Connection) Receive() {
	// create a buffered reader
	buf := bufio.NewReader(connection.socket)

	// loop forever
	for {
		// read the first 4 bytes
		dataSizeBytes := make([]byte,4)
		_, dataSizeReadErr := buf.Read(dataSizeBytes)
		if dataSizeReadErr != nil {
			fmt.Printf("dataSizeReadErr: %s\n", dataSizeReadErr)
			break
		}

		// convert those 4 bytes to 32 bit unsigned int for size of data to follow
		dataSize := binary.BigEndian.Uint32(dataSizeBytes)

		// prepare a buffer to read the data
		payloadBytes := make([]byte, dataSize)
		// read data
		numDataBytes, err := buf.Read(payloadBytes)
		if uint32(numDataBytes)<dataSize || err != nil {
			// report error here
			break
		}
		//fmt.Printf("payloadBytes: %v\n", payloadBytes)

		// Create a decoder and receive a value.
		reply := &Reply{}
		payloadDecoderBuffer := bytes.NewBuffer(payloadBytes)
		payloadDecoder := gob.NewDecoder(payloadDecoderBuffer)

		err = payloadDecoder.Decode(reply)
		if err != nil {
			fmt.Printf("decode: %s\n", err)
		}

		fmt.Printf("Reply %#v\n", reply)

		// deliver reply to calling func
	}
}
