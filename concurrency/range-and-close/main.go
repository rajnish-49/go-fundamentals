package main

import "fmt"

// When you receive from a channel how do you know when the sender is done sending?
// You could receive forever waiting for more values that never come.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //  // sender says — no more values coming
}

// v, ok := <-c
// ok = true  — got a value, channel still open
// ok = false — channel is closed, no more values


// fibonacci is the producer. main is the consumer. Channel is the pipe between them.
// The close function is only needed in the producer, never the consumer. 
// Why? Because only the producer can know when there are no more values to send. 
// If the consumer could close the channel, it would be possible for the producer to send on a closed channel, which would cause a panic.
func main() {
	c := make(chan int, 10) // buffered channel holds 10 values
	go fibonacci(cap(c), c) // cap(c) = 10, generate 10 fibonacci numbers
	for i := range c {      // receive until channel is closed
		fmt.Println(i)
	}
}
