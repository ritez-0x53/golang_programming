package main

import (
	"fmt"
	"sync"
)

// processChan receives a message from the string channel and prints it
// wg.Done() tells the WaitGroup that this goroutine is finished
func processChan(messageChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("processing message ->", <-messageChan)
}

// processNum receives a number from the int16 channel and prints it
// wg.Done() tells the WaitGroup that this goroutine is finished
func processNum(numChan chan int16, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing number ->", <-numChan)
}

func main() {
	// Create two channels: one for string messages, one for int16 numbers
	messageChan := make(chan string)
	numChan := make(chan int16)

	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// We will launch 2 goroutines, so we add 2 to the WaitGroup counter
	wg.Add(2)

	// Start the goroutines
	go processChan(messageChan, &wg)
	go processNum(numChan, &wg)

	// Send data to the channels (this will unblock the goroutines)
	messageChan <- "ping"
	numChan <- 34

	// Wait until both goroutines finish before exiting main
	wg.Wait()
}
