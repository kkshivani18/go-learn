package main

import (
	"fmt"
	"goBasics/basics/imports"
)

func main() {
	fmt.Println("hello")

	newTicket := imports.Ticket{
		ID:    123,
		Event: "Go Course",
	}

	newTicket.PrintEvent()
	fmt.Println(newTicket)
}
