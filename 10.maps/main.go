package main

import "fmt"

func main() {

	//* creating map

	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "backend"
	fmt.Println(m)

	// get an element
	fmt.Println(m["name"], m["area"])

	//* IMP: if key does not exist it returns zero-0 value for int value
	n := make(map[string]int)
	n["age"] = 25
	n["price"] = 120
	fmt.Println(n["phone"])

	// deleting keys in map
	delete(n, "price")
	fmt.Println(n)
	// deleting everything from map
	clear(n)
	fmt.Println(n)

	// creating map without make
	myMap := map[string]int{"price":20,"phones":2}
	fmt.Println(myMap)
}
