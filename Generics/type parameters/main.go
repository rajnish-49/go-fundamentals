package main

//The problem generics solve
//Imagine you want a function that finds an element in a slice. Without generics you need a separate function for every type:

import (
	"fmt"
)

// One function works for int, string, float64, or any other type.

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i // found — return index
		}
	}
	return -1 // not found
}

func main() {

	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // 2 — found at index 2

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) // -1 — not found
}

//Go functions can be written to work on multiple types using type parameters.
// The type parameters of a function appear between brackets, before the function's arguments.
