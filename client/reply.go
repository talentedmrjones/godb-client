package client

// Reply is received from server
type Reply struct {
	ID     string
	Status float64
	Result Records
	Error  string
}
