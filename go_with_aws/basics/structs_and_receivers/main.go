package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// func changeName(person Person, newName string) {
// 	fmt.Println("name before:", person.Name)
// 	person.Name = newName
// 	fmt.Println("name after:", person.Name)
// }

func (p *Person) changeName(newName string) {
	fmt.Println("address of copy", &p.Name)

	fmt.Println("name before:", p.Name)
	p.Name = newName
	fmt.Println("name after:", p.Name)
}

func main() {
	// myPerson := Person{
	// 	Name: "Shivani",
	// 	Age:  21,
	// }

	myPerson := NewPerson("Shivani", 21)
	fmt.Println("address of allocated memory", &myPerson.Name)
	myPerson.changeName("Sanaina")
	fmt.Println(myPerson)

	myPerson.Name = "Shreya"
	// v format specifier is basically for interfaces, it can accept any type
	fmt.Printf("this is my person %+v\n", myPerson)

	a := 7
	b := &a
	fmt.Println(a)
	fmt.Println(&a)
	*b = 9
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)

	mySlice := []int{
		1, 2, 3,
	}

	for index, val := range mySlice {
		fmt.Println(index, val)
	}

	fmt.Println(mySlice)
}
