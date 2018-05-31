package main

import "fmt"

func main() {
	var (
		old = "2.035165e+06"
		new float64
	)
	n, err := fmt.Sscanf(old, "%e", &new)
	if err != nil {
		fmt.Println(err.Error())
	} else if 1 != n {
		fmt.Println("n is not one")
	}

	fmt.Println(n)

	fmt.Println(int64(new))
}
