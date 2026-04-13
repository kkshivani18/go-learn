package main

import "fmt"

type Notifier interface {
	Notify(order Order)
}

type EmailNotifier struct{}

func (EmailNotifier) Notify(order Order) {
	fmt.Println("EMAIL:", order.Summary())
}

type SMSNotifier struct{}

func (SMSNotifier) Notify(order Order) {
	fmt.Println("SMS:", order.Summary())
}
