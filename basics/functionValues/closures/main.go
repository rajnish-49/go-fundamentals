package main 


import (
	"fmt"
)

// A closure is a function that remembers the variables around it even after the outer function has finished.

func counter() func() int {
    count := 0                    // this variable lives inside counter
    return func() int {
        count += 1                // inner function uses count
        return count
    }
}

//  counter creates a variable count
// It returns a function that increments and returns count
// The returned function remembers count even after counter finishes	


func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
	c := counter()	

	fmt.Println(c())   // 1
    fmt.Println(c())   // 2
    fmt.Println(c())   // 3

	pos , neg := adder () , adder()

	fmt.Println(pos(1))   // 1
	fmt.Println(pos(2))   // 3
	fmt.Println(neg(-1))  // -1
	fmt.Println(neg(-2))  // -3
	

}