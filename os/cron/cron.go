package cron

type Cron struct {
	entries []*Entry
	add     chan *Entry
}

type Job interface {
	Run()
}

type Entry struct {
	Job Job
}

func New() *Cron {
	return &Cron{
		entries: nil,
		add:     make(chan *Entry),
	}
}

func (c *Cron) AddJob(spec string, cmd Job) {
	entry := &Entry{
		Job: cmd,
	}
	c.add <- entry
}
