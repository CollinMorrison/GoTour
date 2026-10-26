package main

import (
	"fmt"
)

// a struct is a collection of fields, like an object
// this struct is exported, so it can be used within the package
// or by another package if this one is imported (see ../arrays/main.go)
type Vertex struct {
	X int
	Y int
}

func structs() {
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)

	// a struct's fields are accessibly via a '.'
	v.X = 4
	fmt.Println(v.X, v.Y)

	// we can also access struct values through a struct pointer with no dereferencing needed
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// we can also declare grouped variables like this
	var (
		// all of these are valid delarations of Vertex
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implied
		v3 = Vertex{}      // X:0 and Y:0
		p1 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p1, v2, v3)
}
