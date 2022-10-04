package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Printf("aMap of type %T and value : %v\n", aMap, aMap)
	aMap = nil
	// aMap["test"] = 1 Assigning to a nil map will cause a panic error
	if aMap == nil {
		fmt.Printf("Nil aMap of type %T and value : %v\n", aMap, aMap)
		aMap = map[string]int{}
	}

}
