package main 

import (
	"fmt"
	"math"
)	

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2)) // for comparison
}
