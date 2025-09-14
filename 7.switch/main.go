package main

import (
	"fmt"
	"time"
)

func main() {

	//* simple switch statement

	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Other")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Workdays")
	}

	// type switch
	getType := func(val interface{}) {
		switch val.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other ")
		}
	}
	getType(77.01)

	
}
