package main

import "fmt"

// The important behavior — blocking
// By default a channel blocks. Meaning:
// sending blocks until someone is ready to receive:

//If Person A puts something in the pipe, A waits there until Person B comes to take it out. A cannot walk away.
// If Person B comes to take something out but nothing is there yet, B waits until Person A puts something in. B cannot walk away.


//goroutines run concurrently. But how do they talk to each other? How does one goroutine give a result to another?
//That's what channels are for — they are a pipe between goroutines for passing data.

func main() {
	
	ch := make(chan int)   // a pipe that carries integers

	ch <- 42    // sender STOPS here and waits until someone receives
	v := <-ch   // receiver STOPS here and waits until someone sends
}