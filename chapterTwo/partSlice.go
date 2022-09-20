package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)
	// First 5 elements
	fmt.Println(aSlice[0:5]) // First 5 elements fmt.Println(aSlice[:5])
	// Get first 2 elements
	fmt.Println(aSlice[:2])
	// Get last element
	fmt.Println(aSlice[l-2:])
	// Get elements between 2 and 5 positions
	fmt.Println(aSlice[2:5])
	// Get elements between 2 and 5 positions in a slice with a capacity of 5 - 2 (initial position)
	fmt.Println("L:", len(aSlice[2:5:5]), "C:", cap(aSlice[2:5:5]))
	// Get elements between 1 and 5 positions in a slice with a capacity of 6 - 1 (initial position)
	fmt.Println("L:", len(aSlice[1:5:6]), "C:", cap(aSlice[1:5:6]))

}
