package main

import "fmt"

type Order struct {
	ID       int
	Customer string
	Amount   float64
	Status   OrderStatus
}

func (o *Order) UpdateStatus(newStatus OrderStatus) {
	o.Status = newStatus
}

func (o Order) IsComplete() bool {
	return o.Status == StatusDelivered || o.Status == StatusCancelled
}

func (o Order) Summary() string {
	return fmt.Sprintf(
		"Order %d for %s: %s, $%.2f",
		o.ID,
		o.Customer,
		o.Status,
		o.Amount,
	)
}
