package main

import (
	"fmt"
	"math"
)

/**
An interface defines a shape
A struct implements an interface

Since the Shape interface is stored in a different file but is in this same package,
it is automatically accessible from this file.
That is why there are no errors in the parameters of the calculateArea() function.
**/

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Because the Rectangle struct has an Area() function, it is now a shape
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Because the Circle struct has an Area() function, it is now a shape
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

// capitalized functions are exported and automatically callable within their package
// non-capitalized functions are also callable within their package.
// capitalizing functions exports them, but not sure what that means yet.
// probably makes them callable from outside their package.
func interfaces() {
	fmt.Println("hello imported Interfaces")

	rect := Rectangle{width: 5, height: 4}
	circle := Circle{radius: 2}

	fmt.Println(
		"Rectangle area: ",
		calculateArea(rect),
		"Circle area: ",
		calculateArea(circle))
}
