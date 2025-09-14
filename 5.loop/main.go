package main

import "fmt"

func main() {

	//classic for loop
	for i := 0; i < 3; i++ {
		if i == 0 {
			// break
			continue
		}
		fmt.Println(i)
	}

	// go1.22 , range
	for i := range 5 {
		fmt.Println(i)
	}

}
