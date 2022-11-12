package main

import "fmt"

type IntA interface {
	foo()
}
type IntB interface {
	bar()
}
type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

type compose struct {
	field1 int
	a
}

func (A a) A() {
	fmt.Println("Function A() for A")
}
func (B b) A() {
	fmt.Println("Function A() for B")
}
func main() {
	var iC c = c{a{120, 12}, b{"-12", -12}}
	iC.A.A()
	iC.B.A()
	// The following will not work
	// iComp := compose{field1: 123, a{456, 789}}
	// iComp := compose{field1: 123, XX: 456, YY: 789}
	iComp := compose{123, a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)
	iC.bar()
	processA(iC)
}
