package main

import (
	"fmt"
	"math"
)

// a method is just a function with a special receiver argument — it's attached to a type.
// The receiver appears in its own argument list between the func keyword and the method name.

type Vertex struct {
	X, Y float64
}

// Abs method has a receiver type of Vertex, which means it can be called using a Vertex value.
// Abs is the method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Here's Abs written as a regular function with no change in functionality.

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can attach methods to any type, not just structs
type MyInt int

func (m MyInt) double() MyInt {
	return m * 2
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	j := Vertex{3, 4}
	fmt.Println(Abs(j))

	x := MyInt(5)
	fmt.Println(x.double()) // 10
}
