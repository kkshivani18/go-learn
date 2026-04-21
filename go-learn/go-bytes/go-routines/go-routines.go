package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from go-routine")
}

func worker() {
	fmt.Println("worker running")
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	go worker()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello from main")
}
