package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file:", file)
	path, err := filepath.Abs(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("path:", path)
	base := filepath.Base(path)
	fmt.Println("base:", base)
	dir := filepath.Dir(path)
	fmt.Println("dir:", dir)
	path, err = GetCurrentPath()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
