package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func task(key string, out chan<- status, times int) {
	out <- status{
		key, 0, "start", "",
	}
	time.Sleep(1 * time.Second)
	for i := 1; i <= times; i++ {
		time.Sleep(time.Second)
		out <- status{
			key, i, "running", "",
		}
	}
	time.Sleep(1 * time.Second)
	out <- status{
		key, times, "done", "",
	}
}

type CRONItem struct {
	Schedule   string   `json:"schedule"`
	Command    string   `json:"command"`
	Parameters []string `json:"parameters"`
	Last       string
	Status     bool
	Error      string
	Progress   string
}

func (c *CRONItem) Run() {
	t := time.Now()
	log.Println("[INFO] - Module:[report] -> Running task:", c.Command)
	cmd := exec.Command(c.Command)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	err = cmd.Start()
	defer func() {
		if err != nil {
			log.Println("[ERR] - Module:[report] -> :", err)
			c.Status = false
			c.Error = fmt.Sprint(err)
		} else {
			c.Status = true
			c.Error = ""
		}
		c.Last = time.Since(t).String()
	}()
	if err != nil {
		return
	}
	input := bufio.NewScanner(stdout)
	for input.Scan() {
		c.Progress = input.Text()
	}

	errContent, err := ioutil.ReadAll(stderr)
	if err != nil {
		return
	}
	if err = cmd.Wait(); err != nil {
		err = fmt.Errorf(" %v: %s", err, errContent)
		return
	}
}

func setupCorns() {
	for _, item := range CRONs {
		func(item *CRONItem) {
			schedule := item.Schedule
			cronrunner.AddJob(schedule, item)
		}(item)
	}
	cronrunner.Start()
}

func start(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		return err
	}
	input := bufio.NewScanner(stdout)
	for input.Scan() {
		input.Text()
	}

	errContent, err := ioutil.ReadAll(stderr)
	if err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf(" %v: %s", err, errContent)
	}
	return nil
}
