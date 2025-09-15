package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func NewOrder(id string, amount float32, status string, createdAt time.Time) *order {

	neworder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}
	return &neworder
}

func main() {

	order1 := order{
		id:     "1",
		amount: 2000,
		status: "dispatched",
	}
	order1.createdAt = time.Now()

	fmt.Println(order1)
	fmt.Println(order1.createdAt)

	order2 := order{
		id:        "2",
		amount:    400,
		status:    "received",
		createdAt: time.Now(),
	}

	order2.changeStatus("delivered")
	fmt.Println(order2.status)
	fmt.Println(order2.getAmount())

	myOrderx := NewOrder("3", 2000, "delivered", time.Now())
	myOrderx.changeStatus("completed")
	fmt.Println(myOrderx.status)

	// shorthand
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
