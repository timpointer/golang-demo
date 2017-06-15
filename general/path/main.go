package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)
	fmt.Println(exPath)
}
