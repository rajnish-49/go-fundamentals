package main

import (
	"fmt"
)

func getFunction() func() {
	return func() {
		fmt.Println("hello")
	}
}

func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor    // using factor from outer function
    }
}

func main() {
	x := getFunction() //  x is a function
	x()                // now you call that function → prints "hello"

	double := multiplier(2)   // returns a function with factor=2 baked in
    fmt.Println(double(5))    // 5 * 2 = 10
    fmt.Println(double(10))   // 10 * 2 = 20

    triple := multiplier(3)   // returns a function with factor=3 baked in
    fmt.Println(triple(5))    // 5 * 3 = 15
	
}
