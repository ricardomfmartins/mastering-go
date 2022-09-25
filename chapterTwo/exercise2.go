package main

import (
	"fmt"
	"sort"
)

func main() {

	inputArray1 := [8]int{0, 1, -2, 3, -4, 5, 6, 7}
	fmt.Printf("Type of inputArray1: %T\n", inputArray1)
	inputArray2 := [6]int{-1, 2, -3, 4, -5, -6}
	fmt.Printf("Type of inputArray2: %T\n", inputArray2)
	const outputArrayLength int = len(inputArray1) + len(inputArray2)
	var outputArray [outputArrayLength]int
	var tempSlice []int = inputArray1[:]
	tempSlice = append(tempSlice, inputArray2[:]...)
	sort.Sort(sort.IntSlice(tempSlice))
	copy(outputArray[:], tempSlice)
	fmt.Printf("outputArray as the type %T and the value %v \n", outputArray, outputArray)
}
