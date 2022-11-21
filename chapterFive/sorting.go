package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func newFunction(isSorted bool) {
	if isSorted {
		fmt.Println("It is sorted")
	} else {
		fmt.Println("It is NOT sorted")
	}
}

func main() {
	data := []Grades{{"J.", "Lewis", 9}, {"M.", "Tsoukalos", 12}, {"D.", "Tsoukalos", 10}}

	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	newFunction(isSorted)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})
	newFunction(isSorted)
	fmt.Println("List by Grade:", data)
}
