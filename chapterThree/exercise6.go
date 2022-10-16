package main

import (
	"fmt"
	"os"
)

type Argument struct {
	index int
	value string
}

var args []Argument

func main() {
	args = []Argument{}
	for k, v := range os.Args[1:] {
		temp := Argument{k, v}
		args = append(args, temp)
	}
	fmt.Println(args)

}
