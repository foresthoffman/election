package bully

import "time"

// Leader describes a Voter that has been elected by its peers.
type Leader struct {
	Id        interface{}   `json:"id"`
	Term      time.Duration `json:"term"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}

type Response struct {
	Id     interface{} `json:"id"`
	Leader *Leader     `json:"leader,omitempty"`
}
