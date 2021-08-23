package bully

type Forum interface {
	// Open establishes a connection with a remote platform, through which voters
	// may communicate.
	Open() error
	// IsOpen return the current state of the connection.
	IsOpen() bool
	// Send sends a message to the provided connection.
	Send(i interface{}) error
	// Close the connection.
	Close() error
}

type ElectMessage struct {
	Id int64
}
