// // Control Structures

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Conditional Statements

	age := 30

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// Switches

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("It's the weekend")
	}

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// Loops

	for i := 0; i < 5; i++ {
		fmt.Println("this is i:", i, "ohh")
	}

	// Go doesn't have while loop, you simply use for loop
	counter := 0
	for counter < 3 {
		fmt.Println("counter:", counter)
		counter++
	}

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// Arrays and Slices

	numbers := [5]int{10, 20, 30, 40, 50}
	// numbers := [...]int{10, 20, 30, 40, 50}   // same as above
	numbers[0] = 89

	fmt.Printf("number is: %v\n", numbers)
	fmt.Printf("last val is: %v\n", numbers[len(numbers)-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("this is the matrix: %v\n", matrix)

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// Slices
	// Slices are dynamic Arrays. They are not bound by a predefined size and can grow as more items are inserted.

	// numbers := [5]int{10, 20, 30, 40, 50}

	allNumbers := numbers[:]
	firstThree := numbers[0:3]

	fmt.Println("allNums: ", allNumbers, "firstThree: ", firstThree)

	fruits := []string{"apple", "banana", "strawberry"}
	fmt.Printf("these are fruits: %v\n", fruits)

	fruits = append(fruits, "kiwi", "persimmon")
	fmt.Printf("these are fruits: %v\n", fruits)

	moreFruits := []string{"blueberry", "blackberry"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are fruits: %v\n", fruits)

	for index, value := range numbers {
		fmt.Printf("index %d and value %d\n", index, value)
	}

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// Maps
	// used for key value pairs, hashmaps

	capitalCities := map[string]string{
		"India":   "New Delhi",
		"Finland": "Helsinki",
		"USA":     "Washington D.C.",
	}

	fmt.Println(capitalCities["Finland"])
	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("this is the capital", capital)
	} else {
		fmt.Println("Doesn't exist")
	}

	delete(capitalCities, "USA")
	fmt.Printf("this is new capitalCities %v\n", capitalCities)

	//------------------------------------------------------------------------------------------------------------------------------------------

	// Structs, Struct Methods and Pointers
	// They are data types that can hold data and be passed around the application.
	// The syntax for a struct resembles a class or type declaration from OOP languages.

	// When a struct is defined. type Address struct --> here Address (A) capital means public (a) means private

	person := Person{Name: "John", Age: 25}
	fmt.Printf("This is person %v\n", person)

	employee := struct {
		name string
		id   int
	}{
		name: "Alice",
		id:   123,
	}
	fmt.Printf("This is employee %v\n", employee)

	type Address struct {
		Street string
		City   string
	}

	type Contract struct {
		Name    string
		Address Address
		Phone   string
	}

	contract := Contract{
		Name: "Marc",
		Address: Address{
			Street: "123 Main street",
			City:   "Anytown",
		},
	}
	fmt.Println("this is contract:", contract)

	fmt.Println("name before: ", person.Name) // prints John
	// modifyPersonName("Rosie")
	person.modifyPersonName("Rosie")
	fmt.Println("name after: ", person.Name) // prints John
	// after using &person, above prints Nina

	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 40
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)

	//-------------------------------------------------------------------------------------------------------------------------------------------

	// JSON Struct Example
	// Define a struct with JSON tags

	type User struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Age     int    `json:"age"`
		Active  bool   `json:"active"`
		Address struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}

	user := User{
		Name:   "Jane Doe",
		Email:  "jane@example.com",
		Age:    28,
		Active: true,
	}
	user.Address.Street = "456 Oak Ave"
	user.Address.City = "Springfield"

	// Marshal struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Printf("JSON output: %s\n", jsonData)

	// Marshal with indentation (pretty print)
	jsonDataPretty, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Printf("Pretty JSON output:\n%s\n", jsonDataPretty)
}

type Person struct {
	Name string
	Age  int
}

// func modifyPersonName(person *Person) {
// 	person.Name = "Nina"
// 	fmt.Println("inside scope: new name", person.Name) // prints Nina
// }

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope: new name", p.Name) // prints Nina
}
