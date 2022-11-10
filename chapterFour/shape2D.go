package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	a := circle{R: 15}
	fmt.Printf("R %.2F -> Perimeter %.3f \n", a.R, a.Perimeter())

	interf, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
	fmt.Println(interf, interf.Perimeter())
	fmt.Println(Shape2D(a), Shape2D(a).Perimeter())
}
