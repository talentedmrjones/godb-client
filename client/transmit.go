package client

import (
	"bytes"
	"encoding/gob"
	"encoding/binary"
	"fmt"
)

// decoupling this from command so we can make use of connection pooling later
func transmit (patient *Patient) *Reply {

	command := NewCommand(patient.action, patient)
	patient.table.db.connection.replies[command.Id] = make(chan *Reply)

	// gob encode payload
	var payloadEncodingBuffer bytes.Buffer
	payloadEncoder := gob.NewEncoder(&payloadEncodingBuffer)
	payloadEncodingErr := payloadEncoder.Encode(command)
	if payloadEncodingErr != nil {
		// TODO handle error
		fmt.Printf("%v", payloadEncodingErr)
	}
	payloadBytes := payloadEncodingBuffer.Bytes()

	//fmt.Printf("%v %v", encodeErr, payloadBytes)
	dataSize := make([]byte,4)
	binary.BigEndian.PutUint32(dataSize, uint32(len(payloadBytes)))

	patient.table.db.connection.socket.Write(dataSize)
	patient.table.db.connection.socket.Write(payloadBytes)

	reply := <- patient.table.db.connection.replies[command.Id]
	return reply
}
