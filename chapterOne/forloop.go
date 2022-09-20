package main

import (
	"fmt"
)

var Global int = 1234
var AnotherGlobal = -5678

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i*i, " ")
	}
	fmt.Println()
	a := 0
	for ok := true; ok; ok = (a != 10) {
		fmt.Println(a*a, " ")
		a++
	}
	fmt.Println()
	b := 0
	for {
		if b == 10 {
			break
		}
		fmt.Println(b*b, " ")
		b++
	}
	fmt.Println()
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for k, v := range aSlice {
		fmt.Println("index:", k, "square_value:", v*v)
	}
}
