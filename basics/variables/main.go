package main

import "fmt"

var c, python, java bool

func variable() {
	var i int
	fmt.Println(i, c, python, java)
}

func main() {
	variable()
	shortVariableDemo()
	typesDemo()
}
