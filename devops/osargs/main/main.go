package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	flag.StringVar(&name, "name", name, "user name")
	flag.Parse()
	fmt.Println("**********args**********")
	fmt.Println(strings.Join(os.Args, "\n"))
	fmt.Println("**********params**********")
	fmt.Println("name:", name)
}
