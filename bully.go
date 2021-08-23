// Package bully defines generic patterns that can be used to create a
// leader-election system for distributed computing (*). "Leader-election" or the
// "Bully algorithm" describes the behavior of a system in which N participants
// ([]bully.Voter) interact with a shared medium (bully.Forum) to determine which
// among them is most fit for being in charge. A dynamically assigned, but static
// and unique, identifier is used to determine the fitness of the each Voter.
// Each Voter must ensure they are registered with the shared Forum and send all
// the others a message indicating its intention to be elected. The Voter with
// the highest unique identifier becomes the leader and sends a message to all
// others on the Forum to indicate that it won.
//
// The implementations provided offer one path for implementing this on your own,
// however it isn't recommended to mix and match bully.Voter instances, as they
// could potentially act differently. The name of the game is redundancy
// and stability. The more Voters are added to the pool, the more potential for
// stability. Keep in mind that for certain applications of the leader-election
// system, ensuring bully.Forum redundancy can be just as important.
//
// (*) References can be found here: https://en.wikipedia.org/wiki/Bully_algorithm
package bully

import "time"

func NewTimestampId() int64 {
	return time.Now().UnixNano()
}

var _ Voter = &Bully{}

type Bully struct {
	// Id maintains the current identifier of the Bully. This should not change after
	// the first registration request, unless there was a conflict with another
	// bully's identifier.
	id int64
	// isLeader maintains the current election state of the Bully.
	isLeader bool
	// term indicates the time to live for the Bully. Each time the Bully makes a
	// Ping request, its timer is reset on the Forum.
	term  time.Duration
	forum Forum
}

func (b *Bully) Elect() error {
	panic("implement me")
}

func (b *Bully) Ping() error {
	panic("implement me")
}

func (b *Bully) Register() error {
	panic("implement me")
}

func (b *Bully) Id() interface{} {
	panic("implement me")
}

func (b *Bully) SetId(i interface{}) error {
	panic("implement me")
}

func (b *Bully) Term() time.Duration {
	panic("implement me")
}

func (b *Bully) SetTerm(d time.Duration) error {
	panic("implement me")
}

func (b *Bully) Forum() Forum {
	panic("implement me")
}

func (b *Bully) SetForum(f Forum) error {
	panic("implement me")
}
