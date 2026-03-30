package main

// The problem
// You already know a struct holds data:
// type Node struct {
//     next *Node
//     val  int    // can only hold int
// }

import (
	"fmt"
)

// But this only works for int. If you want a linked list of strings you need to write a whole new struct. 
// Same problem as before with functions.

// Generic type solves this

type List[T any] struct {
    next *List[T]
    val  T
}


func main () {

	// linked list of ints
head := List[int]{val: 10}
head.next = &List[int]{val: 20}
head.next.next = &List[int]{val: 30}

// linked list of strings
head2 := List[string]{val: "hello"}
head2.next = &List[string]{val: "world"}
	
}