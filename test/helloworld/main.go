package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	log.Println("er44")
	fmt.Println("hello world")
	fmt.Println("Hello %q", os.Args[0])
	dir, _ := path.Split("/wer/wer/wer/wet")
	fmt.Println("Hello %q", dir)
	fmt.Println("Hello %q", path.Dir("/wer/wer/wer/wet"))
}
