package main

import (
	"fmt"
)

func main() {
	/* IF STATEMENTS */

	i := 1

	if i > 4 {
		fmt.Println("i > 2");
	} else if i == 3 {
		fmt.Println("i == 3");
	} else {
		fmt.Println("i is less than 3")
	}

	// Define a variable just for this condition
	if num := 9; num < 0 {
		fmt.Println(num, "(num) is negative")
	}
	
}
