package election

import "errors"

var (
	ErrInvalidForum = errors.New("invalid forum")
	ErrClosedForum  = errors.New("forum was closed")
	ErrSendMessage  = errors.New("failed to send message")
)

type Forum interface {
	// Open establishes a connection with a remote platform, through which voters
	// may communicate.
	Open() error
	// IsOpen return the current state of the connection.
	IsOpen() bool
	// Send sends a message to the provided connection.
	Send(i interface{}) ([]byte, error)
	// Close the connection.
	Close() error
}
