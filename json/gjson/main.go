package main

import (
	"encoding/json"
	"fmt"

	"net/url"

	"os"

	"github.com/tidwall/gjson"
)

func main() {
	values := make(url.Values)
	values.Add("name", "tim")
	values.Add("age", "34")
	values.Add("name", "timweg")
	values.Set("age", "geg")
	values.Set("color", "geg")
	js, err := json.Marshal(values)
	if err != nil {
		return
	}
	fmt.Println(string(js))

	arg := os.Args[1]
	fmt.Println(arg)
	name := gjson.Get(arg, "name")
	fmt.Println("*********")
	fmt.Println(name)
	fmt.Println("*********")
}
