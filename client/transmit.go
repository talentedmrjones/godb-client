package client

import (
	"bytes"
	"encoding/gob"
	"encoding/binary"
	"fmt"
	"net"
)

// decoupling this from command so we can make use of connection pooling later
func transmit (conn net.Conn, command *Command) {

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

	conn.Write(dataSize)
	conn.Write(payloadBytes)
}
