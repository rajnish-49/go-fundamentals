package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// walkHelper recursively traverses the tree in-order (left → current → right)
// and sends each value into the channel
func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return // base case — empty node, stop recursion
	}
	walkHelper(t.Left, ch)  // go left first
	ch <- t.Value           // send current node value
	walkHelper(t.Right, ch) // then go right
}

// Walk is the public entry point — starts the traversal and closes
// the channel when all values have been sent
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch) // signal to receiver that no more values are coming
}

// Same returns true if t1 and t2 contain the same values in the same order
// it walks both trees concurrently and compares values as they come out
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1) // walk t1 in background, sends sorted values into ch1
	go Walk(t2, ch2) // walk t2 in background, sends sorted values into ch2

	for v1 := range ch1 {
		v2 := <-ch2      // receive corresponding value from t2
		if v1 != v2 {
			return false // mismatch found — trees are different
		}
	}
	return true // all values matched
}

func main() {
	// test Walk — should print 1 2 3 4 5 6 7 8 9 10 in order
	ch := make(chan int)
	go Walk(tree.New(1), ch) // tree.New(1) builds a sorted tree with 1,2,...,10
	for v := range ch {
		fmt.Println(v) // range stops automatically when Walk closes ch
	}

	// test Same
	fmt.Println(Same(tree.New(1), tree.New(1))) // true  — same values 1..10
	fmt.Println(Same(tree.New(1), tree.New(2))) // false — different values
}