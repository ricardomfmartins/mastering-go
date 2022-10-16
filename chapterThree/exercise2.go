package main

import "fmt"

func main() {
	myArray := [8]string{"Hello!", "This", "is", "a", "sentence", "with", "eigth", "words."}
	myMap := make(map[int]string)
	for k, v := range myArray {
		myMap[k] = v
	}
	fmt.Println("Mymap:", myMap)

	myIndexSlice := []int{}
	myValueSlice := []string{}

	for k, v := range myMap {
		myIndexSlice = append(myIndexSlice, k)
		myValueSlice = append(myValueSlice, v)
	}
	fmt.Println("Index slice:", myIndexSlice)
	fmt.Println("Value slice:", myValueSlice)
	fmt.Println("Values at map:", myMap)
	fmt.Println("Values at index", myIndexSlice[3], ":", myValueSlice[3])
}
