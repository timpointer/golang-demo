package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func command(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return content, nil
}

func start(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return err
}

func hello(name string) string {
	return name
}
