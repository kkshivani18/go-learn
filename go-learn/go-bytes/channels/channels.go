package main

import "fmt"

func sendNumbers(numbers chan int) {
	numbers <- 100
}

func sendMessages(messages chan string) {
	messages <- "Hello from go-routine"
}

func main() {
	messages := make(chan string)
	numbers := make(chan int)

	go sendMessages(messages)
	go sendNumbers(numbers)

	msg := <-messages
	fmt.Println(msg)

	value := <-numbers
	fmt.Println("Received:", value)
}
