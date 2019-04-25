package main

import (
	"fmt"
)

// What are GoRoutines?
// A goroutine is a lightweight thread of execution (concurrent programming)

func main() {

	f("direct")

	//go f("goroutine")

	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")

	//fmt.Scanln()
	//fmt.Println("done")
	
}

func f(from string) {
	for i := 0; i < 3; i++ {
		//time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}