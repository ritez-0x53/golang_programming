package main

import (
	"fmt"
	"sync"
)

func task(id int16, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("running task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := range 6 {
		wg.Add(1)
		go task(int16(i), &wg)
	}

	wg.Wait()
}
