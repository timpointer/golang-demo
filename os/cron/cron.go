package cron

type Cron struct {
	entries []*Entry
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
	}
}

func (c *Cron) AddJob(spec string, cmd Job) error {

}

func (c *Cron) Schedule(schedule Schedule, cmd Job) {
	entry := &Entry{
		Job: cmd,
	}
}
