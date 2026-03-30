package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	//When say("hello") finishes — main exits — program ends 
	// — say("world") also stops even if it's not done.



	go say("world")
	say("hello")
}
