package main

import "fmt"

func main() {
	// cannot change once defined
	const country = "India"
	// country = "Not allowed"

	// grouping multiple constants
	const (
		port = 2000
		host = "localhost"
	)
	fmt.Println(port, host)

}
