package main

import (
	"fmt"
	"math"
)

//A function can be stored in a variable just like a number or string.
// fn := add       // stores the function itself — no parentheses
// fn := add()     // CALLS the function and stores the result

			// fn is of type func(int, int) int
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) // passes hypot to compute
	// compute calls hypot(3, 4)

	fmt.Println(compute(math.Pow))
	// passes math.Pow to compute
	// compute calls math.Pow(3, 4)
	// 3^4 = 81
}
