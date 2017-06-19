package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	result, err := command("./bin/bin")
	if err != nil {
		log.Fatal("err:", err)
		return
	}
	fmt.Println(string(result))
}

func command(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	errContent, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("%s: %v", errContent, err)
	}
	return content, nil
}
