package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	info := util.GetCouponInfo("templte")
	util.GetCouponInfo("template")
	for index := 0; index < 10; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			util.GetCouponInfo("template")
			util.GetCouponInfo("templte")
		}()
	}
	fmt.Println(info)

	//fmt.Println(f)
}
