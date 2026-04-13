package main

type OrderStore interface {
	Save(order Order)
	Load(id int) (Order, bool)
}

type InMemoryStore struct {
	orders map[int]Order
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		orders: make(map[int]Order),
	}
}

func (s *InMemoryStore) Save(order Order) {
	s.orders[order.ID] = order
}

func (s *InMemoryStore) Load(id int) (Order, bool) {
	order, ok := s.orders[id]
	return order, ok
}
