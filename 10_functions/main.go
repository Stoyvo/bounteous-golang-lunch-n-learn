package main

import (
	"fmt"
	"time"
)

func main() {

	a := getSlice()
	fmt.Println("a:", a)

	year,month,day := getCurrentDate() // match function return count
	fmt.Println(year, month, day)

	sumDynamicIntParams(2, 2)
	sumDynamicIntParams(3, 2)

	fmt.Println(func() string {
		return "Anonymous Function (Closures)"
	}())

}

func getSlice() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func getCurrentDate() (int, string, int) {
	t := time.Now()

	return t.Year(), t.Month().String(), t.Day()
}

func sumDynamicIntParams (nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}