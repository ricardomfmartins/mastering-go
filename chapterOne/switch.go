package main

import (
	"fmt"
	//"math"
	"os"
	"strconv"
)

var Global int = 1234
var AnotherGlobal = -5678

// func main() {
// 	var j int
// 	i := Global + AnotherGlobal
// 	fmt.Println("Initial j value:", j)
// 	j = Global
// 	// math.Abs() requires a float64 parameter
// 	//so we type cast it appropriately
// 	k := math.Abs(float64(AnotherGlobal))
// 	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
// }

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line agrument")
		return
	}
	argument := os.Args[1]
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("other number")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}
	switch {
	case value == 0:
		fmt.Println(("Zero!"))
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}
}
}
