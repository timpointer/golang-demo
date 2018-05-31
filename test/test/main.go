package main

import (
	"log"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if true && print() {

	}
}

func print() bool {
	log.Println("sdfs")
	return true
}
