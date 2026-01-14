package main

import (
	"fmt"
	"slices"
)

func main() {

	// Arrays
	animals := [2]string{}
	animals[0] = "dog"
	animals[1] = "cat"

	fruits := [3]string{
		"apple",
		"mango",
		"cherry",
	}

	fmt.Println(animals)
	fmt.Println(fruits)

	//---------------------------------------------------------------------------

	// Slice

	veggies := []string{
		"califlower",
		"cabbage",
		"broccoli",
	}

	veggies = append(veggies, "okra")
	fmt.Println("before delete", veggies)
	veggies = slices.Delete(veggies, 3, 4)
	fmt.Println("after delete", veggies)

	for i := 0; i < len(veggies); i++ {
		fmt.Printf("this is my veggies %s\n", veggies[i])
	}

	for index, value := range veggies {
		fmt.Printf("this is my index %d and veggies %s\n", index, value)
	}

	// there's no while loop in Go, can use for (when while is needed) (with below given way)

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
