package client

// Reply is received from server
type Reply struct {
	Id			string
	Status  uint16
	Records	[]map[string][]byte
	Error		string
}
