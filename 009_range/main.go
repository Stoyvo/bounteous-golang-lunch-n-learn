package main

import "fmt"

func main() {
	// USING 007_slices tutorial
	a := make([]string, 3)
	fmt.Println("arraya:", a, "| Length:", len(a))

	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	fmt.Println("arraya:", a, "| Length:", len(a))

	for k,v := range a {
		fmt.Println(k, "=>", v)
	}
}
