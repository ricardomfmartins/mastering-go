package main

import (
	"fmt"
	"strconv"
)

type record struct {
	Field1 int
	Field2 string
}

func main() {
	S := []record{}
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		S = append(S, temp)
	}
	fmt.Println("Index 0:", S[0].Field1, S[0].Field2)
	fmt.Println("Number of records:", len(S))
	sum := 0

	for _, r := range S {
		sum += r.Field1
	}
	fmt.Println("Sum:", sum)
}
