package main

import "fmt"

type customer struct {
	name  string
	phone string
}

type order struct {
	itemName string
	price    float32
	status   string
	customer
}

func main() {

	customer1 := customer{
		name:  "rx",
		phone: "323231",
	}

	order1 := order{
		itemName: "shirt",
		price:    2000,
		status:   "dispatched",
		customer: customer1,
	}

	fmt.Println(order1)

}
