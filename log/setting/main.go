package main

import "fmt"

type Point struct {
	a int
	b int
}
type list []string
type value map[string][]string

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	p.ScaleBy(2)
	fmt.Println(p)
	var p2 Point
	fmt.Println(p2)
	var v map[string][]string
	fmt.Println(v)
	fmt.Println(v["23"])
	var s list
	fmt.Println(s)
	s = append(s, "23")
	fmt.Println(s[0])
}

func (p *Point) ScaleBy(s int) {
	p.a = (*p).a * s
	p.b = (*p).b * s
}
