package main

type OrderStatus int

const (
	StatusPending OrderStatus = iota
	StatusPaid
	StatusShipped
	StatusDelivered
	StatusCancelled
)

func (s OrderStatus) String() string {
	return [...]string{"Pending", "Paid", "Shipped", "Delivered", "Cancelled"}[s]
}
