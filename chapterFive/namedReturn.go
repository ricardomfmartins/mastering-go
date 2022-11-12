package main

import (
	"fmt"
	"os"
	"strconv"
)

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return min, max
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Needs 2 integers")
		return
	}
	x, err := strconv.Atoi(os.Args[1])

	y, err2 := strconv.Atoi(os.Args[2])
	if err != nil || err2 != nil {
		return
	}
	min, max := minMax(x, y)
	fmt.Println("min:", min, "; max:", max)
}
