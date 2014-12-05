package client

// Read querys the server
func (query *Query) Find () *Reply {

	query.action = "r"

	reply := transmit(query.Patient)

	return reply

}
