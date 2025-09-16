package main

import (
	"fmt"
)

// sum function takes a channel (res) and two integers (a and b)
func sum(res chan int, a int, b int) {
	add := a + b // add the two numbers
	res <- add   // send the result into the channel
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Running task....")
}

func main() {
	res := make(chan int) // create an unbuffered channel of type int

	go sum(res, 12, 5) // run sum() concurrently in a new goroutine

	final := <-res     // receive (read) the value from the channel (blocking)
	fmt.Println(final) // print the result

	done := make(chan bool)
	go task(done)

	<-done
}
