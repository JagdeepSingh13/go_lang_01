package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	// struct embedding
	customer
}

func newOrder(id string, amount float32, status string) *order {
	myOrd := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrd
}

// receiver-type
// structs automatically de-reference
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmtTax() float32 {
	return o.amount * 0.2
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 20,
		status: "not shipped",
	}
	myOrder.createdAt = time.Now()

	myOrder.changeStatus("Shipped")

	fmt.Println("Amount Tax", myOrder.getAmtTax())

	fmt.Println("order struct", myOrder)

	// structs automatically de-reference
	myOrder2 := newOrder("1", 50, "received")
	fmt.Println(myOrder2)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	myOrder3 := order{
		id:     "3",
		amount: 40,
		status: "unpaid",
		customer: customer{
			name:  "JSingh",
			phone: "123456789",
		},
	}
	fmt.Println(myOrder3)

}
