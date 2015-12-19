package client

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

/*
Connection ...
*/
type Connection struct {
	socket  net.Conn
	records chan Record
	Replies chan *Reply
}

/*
NewConnection ...
*/
func NewConnection(address, port string) *Connection {

	socket, socketErr := net.Dial("tcp", address+":"+port)
	if socketErr != nil {
		log.Fatal(socketErr)
	}

	connection := Connection{socket, make(chan Record), make(chan *Reply)}

	go connection.send()
	go connection.receive()
	return &connection
}

/*
Database ...
*/
func (connection *Connection) Database(dbName string) *Database {
	return &Database{connection, dbName}
}

// Receive continuously looks for data from the socket and relays that to a table's command channel
func (connection *Connection) receive() {
	fmt.Println("connection listening for replies...")
	// create a buffered reader
	buf := bufio.NewReader(connection.socket)
	//defer connection.socket.Close()
	// loop forever

	for {

		// read the first 4 bytes which will represent a uint32 for the size of the data
		dataSizeBytes := make([]byte, 4)
		_, dataSizeReadErr := buf.Read(dataSizeBytes)
		if dataSizeReadErr != nil {
			fmt.Printf("dataSizeReadErr: %s\n", dataSizeReadErr)
			break
		}

		// convert those 4 bytes to uint32 for size of data to follow
		dataSize := binary.BigEndian.Uint32(dataSizeBytes)

		// prepare a buffer to read the data
		payloadBytes := make([]byte, dataSize)
		// read a number of bytes equal to the size of the buffer
		numDataBytes, err := buf.Read(payloadBytes)
		if uint32(numDataBytes) < dataSize || err != nil {
			// report error here
			log.Fatal("wrong data size")
			break
		}
		//fmt.Printf("payloadBytes: %v\n", payloadBytes)

		// Create a decoder and receive a value.
		reply := &Reply{}

		payloadBytesUnmarshalErr := json.Unmarshal(payloadBytes, reply)
		if payloadBytesUnmarshalErr != nil {
			log.Fatal("payloadBytesUnmarshalErr:", payloadBytesUnmarshalErr)
			// TODO: report error to connection
		}

		connection.Replies <- reply

	}
}

// send is run in its own goroutine. It continuously loops over replies channel handling replies for that connection
func (connection *Connection) send() {

	for record := range connection.records {
		command := NewCommand(record)

		// encode command as JSON string
		payloadBytes, commandMarshalErr := json.Marshal(command)
		if commandMarshalErr != nil {
			fmt.Printf("commandMarshalError %v", commandMarshalErr)
		}

		// send JSON
		connection.socket.Write(payloadBytes)

	}
}

// Close ...
func (connection *Connection) Close() {
	close(connection.Replies)
	close(connection.records)
	connection.socket.Close()
}
