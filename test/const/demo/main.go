package main

import (
	"crypto/md5"
	"fmt"
	"runtime"
	"time"
)

const (
	FlagUp = iota
	FlagBroadCast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println("runtime", runtime.NumGoroutine())
	go func() {
		time.Sleep(time.Second)
	}()
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadCast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)
	fmt.Println("runtime", runtime.NumGoroutine())
}

func getScanCodePath(epc string) string {
	category := fmt.Sprintf("%x", md5.Sum([]byte(epc)))
	path := "public/scancode/" + category[0:2] + "/" + epc + ".png"
	return path
}
