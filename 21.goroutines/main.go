package main

import (
	"fmt"
	"time"
)

func Task(id int) {
	fmt.Println("Running task -", id)
}

func main() {

	for i := range 11 {
		// go Task(i)

		go func(i int) {
			fmt.Println(i)
		}(i)

	}

	time.Sleep(time.Second * 2)
}
