package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)

	// Check if key exists
	if _,exists := m["k2"]; !exists {
		fmt.Println("k2 Does Not Exist!")
	}
}
