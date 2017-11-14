package main

import (
	"fmt"
	"path"
)

func main() {
	p := "/d/workspace/projects/golang/src/github.com/timpointer/golang-demo/path/demo"
	fmt.Println("app", path.Dir(p))
}
