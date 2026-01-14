package main

import (
	"fmt"
)

type Server struct {
	Url     string
	IsAlive bool
}

func (s Server) Status() string {
	return fmt.Sprintf("Server %s is alive: %t", s.Url, s.IsAlive)
}

// exercise had write a func, hence, func funcName args
func SetDead(s Server) {
	s.IsAlive = false
}

func SetDeadPtr(s *Server) {
	s.IsAlive = false
}

func main() {
	// fmt.Println("Using exercise 2 in other directory to avoid main() conflict. \nSince, one directory can have main func. \nAs all .go files in the same directory are part of the same package")

	// ex1
	s1 := Server{Url: "http://localhost:8081", IsAlive: true}
	fmt.Println(s1.Status())

	// ex2
	// pass by Value. Creates a photocopy of your struct
	SetDead(s1)
	fmt.Println("After SetDead (No Pointer):", s1.IsAlive)

	// pass by Reference. Gets memory address of the original struct
	SetDeadPtr(&s1)
	fmt.Println("After SetDead (With Pointer):", s1.IsAlive)

	// ex3
	serverPool := []*Server{
		{Url: "http://localhost:8081", IsAlive: true},
		{Url: "http://localhost:8082", IsAlive: false},
		{Url: "http://localhost:8083", IsAlive: true},
		{Url: "http://localhost:8084", IsAlive: false},
	}

	serverPool[3].IsAlive = true

	for _, server := range serverPool {
		if server.IsAlive == true {
			fmt.Println("Forwarding traffic to", server.Url, &server.Url)
		} else {
			fmt.Println("Skipping", server.Url, "(Dead)")
		}
	}
}
