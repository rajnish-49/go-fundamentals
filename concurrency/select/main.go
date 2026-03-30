package main

import "fmt"


//channel receive blocks until someone sends. But what if you are waiting on multiple channels at the same time? 
// How do you know which one will have data first?
// That's what select does — it waits on multiple channels and runs whichever one is ready first.

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}



// Select is like a switch but for channels
// It lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. 
// It chooses one at random if multiple are ready.

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}


