package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return 2 * k
		}
	}

	return func(k int) int {
		return k * k
	}
}

func main() {
	n := 10
	i := funRet(n)
	j := funRet(-4)
	fmt.Printf("%T\n", i)
	//It appears a warning because the function isn't actually called
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))
	fmt.Println(i(10))
	fmt.Println(j(10))
}
