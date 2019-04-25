package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

func main() {
	// USING 12_goroutines
	f("direct") // Main thread

	done := make(chan bool)
	go func() {
		f("goroutine")
		done <- true
	}() // New Thread


	<- done
}

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}