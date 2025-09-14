package main

import "fmt"

func main() {

	age := 11
	if age >= 18 {
		fmt.Print("person is an adult")
	} else {
		fmt.Println("person is a teenage or kid")
	}

	//* We can declare a variable inside if construct
	if age := 13; age >= 18 {
		fmt.Println("adult")
	} else if age >= 12 {
		fmt.Println("teenage")
	} else {
		fmt.Println("kid")
	}

	//* Go doesnot have ternary operator

}
