package main

import (
	"fmt"
	"sort"
)

func main() {

	inputSlice1 := []int{0, 1, -2, 3, -4, 5, 6, 7}
	fmt.Printf("Type of inputSlice1: %T\n", inputSlice1)
	inputSlice2 := []int{-1, 2, -3, 4, -5, -6}
	var tempSlice []int = inputSlice1[:]
	fmt.Printf("Type of inputSlice2: %T\n", inputSlice2)
	tempSlice = append(tempSlice, inputSlice2[:]...)
	sort.Sort(sort.IntSlice(tempSlice))

	// There is no way to dynamically set the array limit on this case,
	// because the calculation of the len of the slice can never return a constant
	// due to the dynamic length of the slice.

	var outputArray [12]int
	copy(outputArray[:], tempSlice)
	fmt.Printf("outputArray as the type %T and the value %v \n", outputArray, outputArray)
}
