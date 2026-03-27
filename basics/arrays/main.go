package main

import "fmt"

// go is a statically typed language, which means that you have to declare the type of a variable when you create it.
//  The length of an array is part of its type, so arrays cannot be resized.

func main() {
	var a [5]int
	fmt.Println("emp:", a)	
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
}