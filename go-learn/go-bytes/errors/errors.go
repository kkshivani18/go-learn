// Write a function called squareRoot that takes an integer and returns:

// The square root (use float64 for the result)
// An error if the input is negative
// Hint: You'll need to import the math package and use math.Sqrt() to calculate the square root. You'll also need to convert the integer input to float64 before passing it to math.Sqrt().

// In main(), call the function with both valid and invalid inputs to test your code. Print the error when one occurs. Print the result when everything is fine.

package main

import (
	"errors"
	"fmt"
	"math"
)

func squareRoot(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot take an input of square root as a negative number")
	}
	return math.Sqrt(float64(n)), nil
}

func main() {
	result, err := squareRoot(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
