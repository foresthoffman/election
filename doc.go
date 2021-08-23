// Package election defines generic patterns that can be used to create a
// leader-election system for distributed computing (*). "Leader-election" or the
// "Bully algorithm" describes the behavior of a system in which N participants
// ([]election.Voter) interact with a shared medium (election.Forum) to determine which
// among them is most fit for being in charge. A dynamically assigned, but static
// and unique, identifier is used to determine the fitness of the each Voter.
// Each Voter must ensure they are registered with the shared Forum and send all
// the others a message indicating its intention to be elected. The Voter with
// the highest unique identifier becomes the leader and sends a message to all
// others on the Forum to indicate that it won.
//
// The implementations provided offer one path for implementing this on your own,
// however it isn't recommended to mix and match election.Voter instances, as they
// could potentially act differently. The name of the game is redundancy
// and stability. The more Voters are added to the pool, the more potential for
// stability. Keep in mind that for certain applications of the leader-election
// system, ensuring election.Forum redundancy can be just as important.
//
// (*) References can be found here: https://en.wikipedia.org/wiki/Bully_algorithm
package election
