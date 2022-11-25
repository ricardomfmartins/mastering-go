package main

import (
	"fmt"
	"os"
	"strconv"
)

func sortThree(x, y, z int) (int, int, int) {
	if x > y {
		if z > x {
			return y, x, z
		} else if y > z {
			return z, y, x
		}
		return y, z, x
	} else {
		//y > x
		if z > y {
			return x, y, z
		} else if x > z {
			return z, x, y
		}
		return x, z, y
	}
}

func namedSortThree(x, y, z int) (min, middle, max int) {
	if x > y {
		if z > x {
			min = y
			middle = x
			max = z
		} else if y > z {
			min = z
			middle = y
			max = x
		} else {
			min = y
			middle = z
			max = x
		}
	} else {
		//y > x
		if z > y {
			min = x
			middle = y
			max = z
		} else if x > z {
			min = z
			middle = x
			max = y
		} else {
			min = x
			middle = z
			max = y
		}
	}
	return min, middle, max
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Needs 3 integers")
		return
	}
	x, err := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	z, err2 := strconv.Atoi(os.Args[3])
	if err != nil || err2 != nil {
		return
	}
	min, middle, max := sortThree(x, y, z)
	fmt.Println("min:", min, "; middle:", middle, "; max:", max)
	min, middle, max = namedSortThree(x, y, z)
	fmt.Println("min:", min, "; middle:", middle, "; max:", max)
}
