package main

import "fmt"

func infloop() {
	sum := 0
	for {
		if sum > 100 {
			break
		}
		sum++
	}
	fmt.Print(sum)
}
