package main

import "fmt"

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderConfirmed
	OrderShipped
	OrderDelivered
	OrderCancelled
)

func (s OrderStatus) String() string {
	return [...]string{
		"Pending",
		"Confirmed",
		"Shipped",
		"Delivered",
		"Cancelled",
	}[s]
}

func main() {
	status := OrderConfirmed

	fmt.Println("Order status:", status)
	fmt.Println("Raw value:", int(status))

	switch status {

	case OrderPending:
		fmt.Println("Waiting for Payment")

	case OrderConfirmed:
		fmt.Println("Payment received")

	case OrderShipped:
		fmt.Println("Package is shipped")

	case OrderDelivered:
		fmt.Println("Order completed")

	case OrderCancelled:
		fmt.Println("Order cancelled")

	default:
		fmt.Println("Unknown status")
	}
}
