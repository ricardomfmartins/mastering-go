package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("An index to delete is required.")
		return
	}

	deleteOpt := 1
	if len(arguments) == 3 {
		option := arguments[2]
		optionInt, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println(err)
			return
		}
		deleteOpt = optionInt
	}
	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
	}
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)
	if i > len(aSlice)-1 {
		fmt.Printf("Index %d out of range\n", i)
		return
	}
	if deleteOpt == 1 {
		aSlice = append(aSlice[:i], aSlice[i+1:]...)
		fmt.Println("First delete option:", aSlice)
	} else {
		aSlice[i] = aSlice[len(aSlice)-1]
		aSlice = aSlice[:len(aSlice)-1]
		fmt.Println("Second delete option:", aSlice)
	}
}
