package main

func main() {
	var done chan string
	var stringStream chan string
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}
}
