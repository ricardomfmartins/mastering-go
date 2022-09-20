package main

import "fmt"

func main() {
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	aSlice = append(aSlice, 1234.45)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	t := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Println(-(i + 1))
		t[i] = -(i + 1)
	}

	fmt.Println(t)

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	makeTwoD := make([][]int, 2)
	fmt.Println(makeTwoD)
	makeTwoD[0] = []int{1, 2, 3, 4}
	makeTwoD[1] = []int{-1, -2, -3, -4}
	fmt.Println(makeTwoD)
}
