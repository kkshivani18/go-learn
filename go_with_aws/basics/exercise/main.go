package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (c Circle) findCircumference() float64 {
	return 2 * PI * c.radius
}

func (c Circle) findArea() float64 {
	return PI * c.radius * c.radius
}

func main() {
	myCircle := Circle{
		radius: 20,
	}

	circleCircumference := myCircle.findCircumference()
	fmt.Printf("Circumference: %.2f\n", circleCircumference)

	circleArea := myCircle.findArea()
	fmt.Printf("Area: %.2f\n", circleArea)
}
