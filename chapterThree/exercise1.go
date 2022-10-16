package main

import "fmt"

func main() {
	myArray := [8]string{"Hello!", "This", "is", "a", "sentence", "with", "eigth", "words."}
	mymap := make(map[int]string)
	for k, v := range myArray {
		mymap[k] = v
	}
	fmt.Println(mymap)
}
