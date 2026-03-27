package main

import "fmt"

//range is just a way to iterate over slices, arrays, maps, strings, and channels — cleaner than a manual for loop
// Skip index or value with _

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	s := []int{10, 20, 30}

	// only value
    for _, v := range s {
    	fmt.Println(v)
	}

	// only index
	for i := range s {
    	fmt.Println(i)
	}	

}
