package main

import (
	"fmt"
)

func Hello() string {
	return "Hello, world"
}

type Person struct {
	Name string
	Age  int
}

func main() {

	// hello world thing
	fmt.Println(Hello())
	fmt.Println("changed")

	//-------------------------------------------------------------------------------------------------------------------------

	// Vars and Basic Types

	var name string = "frontend masters"
	fmt.Printf("I am learning Go from %s\n", name)

	// type inferencing (compiler can infer the type for us)
	age := 22
	fmt.Printf("the age is: %d\n", age)

	// declare first, assign later
	var city string
	city = "Goa"
	fmt.Printf("I wanna go to %s\n", city)

	// declare multiple names
	var country, continent string = "India", "Asia"
	fmt.Printf("I live in %s that's in %s\n", country, continent)

	// declare multiple vars of diff types using parentheses
	var (
		isEmployed bool   = true
		salary     int    = 20000
		position   string = "developer"
	)

	fmt.Printf("isEmployed: %t with salary: %d with position: %s\n", isEmployed, salary, position)

	// zero values
	var defaultInt int
	var defaulFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("defaultInt: %d, defaultFloat: %f, defaultString: %s, defaultBool: %t \n", defaultInt, defaulFloat, defaultString, defaultBool)

	// --------------------------------------------------------------------------------------------------------------------------------------------------

	// Const and Enums

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)
	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d\n", Monday, Tuesday, Wednesday)

	// typed v/s untyped const

	const typedAge int = 25
	const untypedAge = 25
	fmt.Println(typedAge == untypedAge) // true

	const (
		Jan int = iota + 1
		Feb
		Mar
		Apr
	)
	fmt.Printf("Jan: %d, Feb: %d, Mar: %d,  Feb: %d\n", Jan, Feb, Mar, Apr)

	//----------------------------------------------------------------------------------------------------------------------------------------

	// Functions

	result := add(3, 4)
	fmt.Println(result)

	sum, prod := calcSumProd(5, 8)
	fmt.Println(sum, prod)
}

func add(a int, b int) int {
	return a + b
}

func calcSumProd(a, b int) (int, int) {
	return a + b, a * b
}
