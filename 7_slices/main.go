package main

import "fmt"

func main() {
	/* SLICE */
	// Slices have similar markup to arrays, but is a more powerful interface for sequences
	// A slice represents a flexible-length array-like data type while providing full control over memory allocations.

	a := make([]string, 3)
	fmt.Println("arraya:", a, "| Length:", len(a))

	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	fmt.Println("arraya:", a, "| Length:", len(a))

	a = append(a, "d")
	a = append(a, "e")
	fmt.Println("arraya:", a, "| Length:", len(a))
}
