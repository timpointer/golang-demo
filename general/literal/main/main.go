package main

import "fmt"
import "net/url"

type age int

type person struct {
	Name string
	Age  age
}
type args struct {
	values url.Values
}

func main() {

	var values url.Values
	values = map[string][]string{
		"name": []string{"hello"},
	}
	fmt.Println(values)
	var arg args
	arg = args{values}
	fmt.Println(arg)

	myage := []age{age(2)}
	fmt.Println(myage)

	my := person{
		"tim",
		age(2),
	}
	fmt.Println(my)

	mylist := map[string]person{
		"time": person{"tim", age(2)},
	}
	fmt.Println(mylist)

	array := [...]int{1, 2, 3}
	fmt.Println(array)
	arratstr := [...]string{"a", "b", "c"}
	fmt.Println(arratstr)

	list := []int{1, 2, 3}
	fmt.Println(list)

	var map1 = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(map1)

	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
	}
	fmt.Println(prereqs)
}
