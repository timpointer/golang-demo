package mycron

import "gopkg.in/robfig/cron.v2"
import "sync"

type Cron struct {
	*cron.Cron
	sm sync.Mutex
	m  map[string]cron.EntryID
}

type Job interface {
	cron.Job
	ID() string
}

func New() *Cron {
	return &Cron{
		Cron: cron.New(),
		m:    make(map[string]cron.EntryID),
	}
}

// AddJob adds a Job to the Cron to be run on the given schedule.
func (n *Cron) AddJob(spec string, cmd Job) error {
	n.sm.Lock()
	defer n.sm.Unlock()
	id, err := n.Cron.AddJob(spec, cmd)
	n.m[cmd.ID()] = id
	return err
}

// Remove an entry from being run in the future.
func (n *Cron) Remove(ID string) {
	n.sm.Lock()
	defer n.sm.Unlock()
	id := n.m[ID]
	n.Cron.Remove(id)
}
