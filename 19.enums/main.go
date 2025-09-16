package main

import "fmt"

// enumerated types
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Placed
	Delivered
)

type FullForm string

const (
	IND  FullForm = "India"
	US   FullForm = "United States"
	UK   FullForm = "United Kingdom"
	USSR FullForm = "Union of Soviet Socialist Republic"
)

func getOrderStatus(status OrderStatus) {
	fmt.Println("Your Status is", status)
}

func getFullForm(short FullForm) {
	fmt.Println(short)
}

func main() {

	getOrderStatus(Delivered)
	getFullForm(UK)
}
