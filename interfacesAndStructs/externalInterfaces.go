package main

// Capital 'Shape' means this interface is public, and can be imported by external packages
type Shape interface {
	Area() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Measurable interface {
	Perimeter() float64
}

func ExternalInterfaces() {

}
