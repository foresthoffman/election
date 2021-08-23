package bully

import "time"

type Voter interface {
	// Elect sends a message to the forum of all other voters, voting for the voter
	// with the highest id.
	Elect() error
	// Ping sends a message to the forum of all other voters, indicating that it is
	// still alive.
	Ping() error
	// Register sends a message to the forum of all other voters, only upon first
	// start up.
	Register() error
	// Id returns the unique id for the voter.
	Id() interface{}
	// SetId assigns the provided unique id to the voter.
	SetId(i interface{}) error
	// Term returns the time-to-live for the voter. This number may not be the same
	// for all voters, since environments may vary.
	Term() time.Duration
	// SetTerm assigns the provided time-to-live to the voter.
	SetTerm(d time.Duration) error
	// Forum returns a connector that will handle the sending of messages for the voter.
	Forum() Forum
	// SetForum assigns the provided connector, for sending messages, to the voter.
	SetForum(f Forum) error
}
