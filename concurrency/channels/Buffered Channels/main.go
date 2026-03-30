package main 

import (
	"fmt"
)

// Buffered Channels has a waiting room for messages. 
// It can hold a certain number of messages before it starts blocking the sender.

ch := make(chan int, 2)   // buffer size 2 — can hold 2 values


// Think of it like a queue with limited space. Sender can put values in without waiting — as long as the queue is not full.

ch <- 1    // queue: [1]     — not full, no blocking
ch <- 2    // queue: [1, 2]  — not full, no blocking
ch <- 3    // queue is full! — blocks until someone receives

func main() {

	ch := make(chan int, 2)  // buffer size 2 — can hold 2 values
	ch <- 1    // queue: [1]     — not full, no blocking
    ch <- 2    // queue: [1, 2]  — not full, no blocking
    ch <- 3    // queue is full! — blocks until someone receives - DEADLOCK 
}
