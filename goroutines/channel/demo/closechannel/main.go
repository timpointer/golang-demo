package main

func main() {
	task := make(chan int)
	close(task)
	close(task)
}
