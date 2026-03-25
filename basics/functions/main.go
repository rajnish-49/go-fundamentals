// main.go
package main

import "fmt" // fmt is a standard library package for formatted I/O , gives access to functions like Println

func main() {
	fmt.Println("Functions demos")
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y := split(17)
	fmt.Println(x, y)
}
