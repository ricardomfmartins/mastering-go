package main

import "fmt"

func main() {
	var outputSlice []int

	inputArray1 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Type of inputArray1: %T\n", inputArray1)
	inputArray2 := [6]int{-1, -2, -3, -4, -5, -6}
	fmt.Printf("Type of inputArray2: %T\n", inputArray2)
	outputSlice = inputArray1[:]
	outputSlice = append(outputSlice, inputArray2[:]...)
	fmt.Printf("outputSlice as the type %T and the value %v \n", outputSlice, outputSlice)
}
