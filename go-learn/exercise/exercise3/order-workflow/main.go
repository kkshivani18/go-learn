package main

import "fmt"

func processOrder(
	order *Order,
	newStatus OrderStatus,
	notifier Notifier,
	store OrderStore,
) {
	order.UpdateStatus(newStatus)
	store.Save(*order)
	notifier.Notify(*order)
}

func main() {
	store := NewInMemoryStore()
	notifier := EmailNotifier{}

	order := Order{
		ID:       1,
		Customer: "Alisa",
		Amount:   3400,
		Status:   StatusPending,
	}

	fmt.Println("Initial order:", order.Summary())
	store.Save(order)

	processOrder(&order, StatusPaid, notifier, store)
	processOrder(&order, StatusShipped, notifier, store)
	processOrder(&order, StatusDelivered, notifier, store)

	fmt.Println("Complete?", order.IsComplete())
	saved, ok := store.Load(order.ID)
	if ok {
		fmt.Println("Loaded from store:", saved.Summary())
	}
}
