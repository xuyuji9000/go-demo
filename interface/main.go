package main

import (
	"fmt"
	"math"
)

/**
- define interface Shape including area() and diameter()
- define type rectangle and circle implement interface Shape
**/

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

// implement area() function from Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// if this function is not implemented, the build will fail
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	circle := Circle{4}
	fmt.Println("Circle radius:", circle.radius)
	fmt.Println("Circle diameter:", circle.perimeter())
	fmt.Println("Circle area:", circle.area())

}
