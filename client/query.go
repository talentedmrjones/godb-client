package client

// Read querys the server
func (query *Query) Find () *Reply {

	query.action = "r"

	reply := transmit(query.Patient)

	return reply

}

// Delete removes a record from the server
func (query *Query) Delete () *Reply {

	query.action = "d"

	reply := transmit(query.Patient)

	return reply

}
