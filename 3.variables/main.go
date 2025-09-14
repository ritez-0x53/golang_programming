package main

import "fmt"

func main() {

	var name string = "ritez"
	// type infering automaticaly
	var age int = 22
	var isAdult bool = true

	fmt.Println(name, age, isAdult)

	// shorthand syntax
	region := "Guwahati"
	fmt.Println(region)

	// first declaring and then defining value
	var language string
	language = "golang"
	fmt.Println(language)

	var price float32
	price = 20.22
	fmt.Println(price)
}
