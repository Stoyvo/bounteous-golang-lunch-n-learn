package main

import (
	"fmt"
)

func main() {
	// Define Variable, then loop. i value increases and stays in function scope
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Define variable for only this loop, j does not exist outside of the loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//fmt.Println(j)

	// Loop forever until break
	for {
		fmt.Println("loop")
		break
	}
	
}
