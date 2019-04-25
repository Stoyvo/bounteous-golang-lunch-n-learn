package main

import "fmt"

/*
 ** Basic Types
 * bool
 * string
 *
 * int  int8  int16  int32  int64
 * uint uint8 uint16 uint32 uint64 uintptr
 *
 * byte // alias for uint8
 *
 * rune // alias for int32 - represents a Unicode code point
 *
 * float32 float64
 *
 * complex64 complex128
 */

const sChallenge string = "try to change me"

func main() {
	// Assigning a variable, assuming type from value
	var a = "initial"
	fmt.Println(a)

	// Assigning a variable, defining int as the type, with values
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Create new variable, assuming type from value
	d := 1
	fmt.Println(d)

	// Uninitialized variable, defining int64 as the type
	var e int64
	fmt.Println(e)

	fmt.Println("const sChallenge: ", sChallenge)
}
