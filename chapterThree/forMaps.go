package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "a value"

	for k, v := range aMap {
		fmt.Println("Key:", k, ",Value:", v)
	}
	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
