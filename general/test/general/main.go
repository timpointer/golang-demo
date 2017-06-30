package main

import "fmt"

func main() {
	m := make(map[string]*int)
	var s int = 2
	m["one"] = &s
	m["two"] = &s
	m["three"] = &s
	l := mToList(m)
	fmt.Println(l)
}

type mytype struct{}

func mToList(m interface{}) interface{} {
	mp, ok := m.(map[string]interface{})
	if ok != true {
		panic("error")
	}
	list := []interface{}{}
	for _, v := range mp {
		list = append(list, v)
	}
	return list
}
