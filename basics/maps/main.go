package main

import "fmt"

// Maps in Go are like unordered_map in C++ — key-value pairs.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

// var m map[string]int — nil map
// No memory allocated for the map yet
// You can read from it safely (returns zero value)
// You cannot write to it — will panic

// m := make(map[string]int) — empty map ready to use
// Now memory is actually allocated
// You can read and write to it safely

func main() {

	m := make(map[string]int)

	// insert / update
	m["alice"] = 25
	m["bob"] = 30

	// read
	fmt.Println(m["alice"]) // 25

	// delete
	delete(m, "alice")

	// length
	fmt.Println(len(m)) // 1


	// Test that a key is present with a two-value assignment:
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
