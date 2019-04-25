package main

import "fmt"

func main() {
	/* ARRAYS */

	var a [5]int // specific length
	fmt.Println("arraya:", a)

	a[2] = 2222
	fmt.Println("arraya:", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrayb:", b)

	var twoD = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("2d: ", twoD)
}
