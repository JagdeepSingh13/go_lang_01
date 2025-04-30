package main

import "fmt"

type orderStatus string

const (
	Received  orderStatus = "Received"
	Confirmed             = "Confirmed"
	Prepared              = "Prepared"
	Delivered             = "Delivered"
)

func changeStatus(status orderStatus) {
	fmt.Println("changing order status to ", status)
}

func main() {
	changeStatus(Delivered)
}
