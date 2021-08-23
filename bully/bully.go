// Package bully is an example implementation of the election package.
package bully

import (
	"encoding/json"
	"fmt"
	"github.com/foresthoffman/election"
	"github.com/foresthoffman/loggy"
	"os"
	"time"
)

var (
	_   election.Voter = &Bully{}
	Log                = loggy.New(os.Stdout, os.Stderr, "", loggy.LevelDebug)
)

type Bully struct {
	// Id maintains the current identifier of the Bully. This should not change after
	// the first registration request, unless there was a conflict with another
	// bully's identifier.
	id int64
	// isLeader maintains the current election state of the Bully.
	isLeader bool
	// term indicates the time to live for the Bully. Each time the Bully makes a
	// Ping request, its timer is reset on the Forum.
	term time.Duration
	// forum maintains a controller for sending messages to adjacent bullies.
	forum election.Forum
}

func (b *Bully) Elect() error {
	return nil
}

func (b *Bully) Ping() error {
	if b.forum == nil {
		return election.ErrInvalidForum
	}
	if !b.forum.IsOpen() {
		return election.ErrClosedForum
	}

	req := Request{Id: b.id}

	Log.Info("sending ping request...")
	Log.Debug(fmt.Sprintf("%v", req))
	body, err := b.forum.Send(req)
	if err != nil {
		err = fmt.Errorf("%s: %w", election.ErrSendMessage.Error(), err)
		Log.Warning(err.Error())
		return err
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response: %w", err)
		Log.Critical(err.Error())
		return err
	}

	if resp.Leader == nil {
		err = b.Elect()
		if err != nil {
			return fmt.Errorf("no leader: %w", err)
		}
	} else if resp.Leader.Id != b.id {
		b.isLeader = false
		err = b.SetTerm(resp.Leader.Term)
		if err != nil {
			err = fmt.Errorf("failed to update leader: %w", err)
			Log.Critical(err.Error())
			return err
		}
	}

	return nil
}

func (b *Bully) Id() interface{} {
	return b.id
}

func (b *Bully) SetId(i interface{}) error {
	id, ok := i.(int64)
	if !ok {
		return fmt.Errorf("failed to cast interface{} to int64: %v", i)
	}

	b.id = id
	return nil
}

func (b *Bully) Term() time.Duration {
	return b.term
}

func (b *Bully) SetTerm(d time.Duration) error {
	if d.Seconds() == 0 {
		return fmt.Errorf("term must be greater than 0s")
	}

	b.term = d
	return nil
}

func (b *Bully) Forum() election.Forum {
	return b.forum
}

func (b *Bully) SetForum(f election.Forum) error {
	if f == nil {
		return election.ErrInvalidForum
	}

	b.forum = f
	return nil
}
