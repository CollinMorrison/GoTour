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

type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

// the source code of Go has an interface 'error' with
// a function Error() that returns a string. So, you can create a struct
// that implements that interface, and is considered an error by Go
func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid input",
		}
	}
	return math.Sqrt(val), nil
}

// Because the Rectangle struct has an Area() function, it is now a shape
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Beacuse the Rectangle struct has Area() and Perimeter() function, it is now a Geometry
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Because the Circle struct has an Area() function, it is now a shape
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
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

	// this line declares variable mysteryBox with value 10,
	// and typecasts it to type interface.
	// An empty interface{} (with no function contract) requires no specific function shape
	// Therefore, any type (int, string, struct, slice, pointer) satisfies the contract
	// An interface value holds dynamic type (%T, int), and dynamic value (%v, 10).
	// This line can also be rewritten as mysteryBox any = 10, because 'any' is
	// an alias for interface{}.
	// So this is a type conversion. 'string(10)' does the same thing, but for strings
	mysteryBox := interface{}(10)
	describeValue(mysteryBox)

	// someValue.(someType) is a type assertion. It checks and extracts an interface's stored value
	// That means this is not valid on any other type - because an empty interface is
	// a type that doesn't mean anything, thus the 'any' alias

	// Thus the following two lines are invalid
	// someInt := 11
	// retrievedInt, ok := someInt.(int)
	retrievedInt, ok := mysteryBox.(int)
	if ok {
		fmt.Println("Retrieved int:", retrievedInt)
	} else {
		fmt.Println("Value is not an integer")
	}

	describeShape(rect)

	val := 4.00
	result, err := performCalculation(val)
	if err != nil {
		fmt.Println("cannot calculate sqrt of", val, "with error:", err)
	} else {
		fmt.Println("sqrt of", val, "is", result)
	}

}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
