package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello arrays!")

	// This declares variable 'a' as an array of type 'string' with 2 spaces
	// You cannot change the size of an array after the fact
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Slices are dynamically sized. Use them instead of arrays in 99% of cases
	// This is a slice - there is not a number in the brackets
	// b creates a slice from a range of elements from the 'prime' array
	// the first index is included, the second one is not.
	// So this slice contains indeces 1, 2, 3 from the 'prime' array
	var b []int = primes[1:4]
	fmt.Println(b)

	// A slice under the hood is actually just an extended array.
	// A slice does not store any data. It just describes a section of an underlying array
	// Therefore, other slices that share the same underlying array will see the changes
	// made by another slice
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// 'first' is a slice of names including indeces 0, 1
	first := names[0:2]
	// 'second' is a slice of names including indeces 1, 2
	second := names[1:3]
	fmt.Println(first, second)

	// the changes to index 0 of second[] also change:
	// index 1 of slice first[]
	// index 1 of array names[]
	second[0] = "X"
	fmt.Println(first, second)
	fmt.Println(names)
}
