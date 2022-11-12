package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 string
	F3 string
}

func diffStructs(s interface{}) {
	switch T := s.(type) {
	case S1:
		fmt.Printf("Data type: S1, %s\n", T)
	case S2:
		fmt.Printf("Data type: S2, %s\n", T)
	default:
		fmt.Println("Unklnown type.")
	}
}

func main() {
	first := S1{
		F1: 123,
		F2: "first",
	}
	second := S2{
		F1: first.F1,
		F2: first.F2,
		F3: "second",
	}
	diffStructs(first)
	diffStructs(second)
}
