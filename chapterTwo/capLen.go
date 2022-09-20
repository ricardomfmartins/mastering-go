package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	// The capacity doubled due to the element being added to the slice
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)
	// The capacity doubled due to the 4 new elements being added to the slice
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))

}
