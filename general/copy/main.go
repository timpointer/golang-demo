package main

import (
	"log"
	"reflect"
	"time"
)

func main() {
	p := &person{Name: "tim", Age: 3}
	log.Println("p", p)
	p2 := *p
	m1 := &man{p}
	m2 := &man{&p2}
	log.Println("m1", m1.person)
	log.Println("m2", m2.person)
	log.Println("m1", m1)
	log.Println("m2", m2)
	ok := reflect.DeepEqual(m1, m2)
	log.Println("ok", ok)
	ok2 := m1.compare(m2)
	log.Println("ok2", ok2)
}

type person struct {
	Name string
	Age  int
	Time time.Time
}

type man struct {
	*person
}

func (m *person) compare(n interface{}) bool {
	return reflect.DeepEqual(m, n)
}
