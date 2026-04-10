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
	names := [...]string{"Pending", "Paid", "Shipped", "Delivered", "Cancelled"}

	if int(s) < 0 || int(s) >= len(names) {
		return "Unknown"
	}

	return names[s]
}
