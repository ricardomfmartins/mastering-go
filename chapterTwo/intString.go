package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Write a number: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("You need to provide a number.")
		return
	}
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)
	input = string(n)
	fmt.Printf("string() %s of type %T\n", input, input)

}
