package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Your last name: ")
	var lastName string
	fmt.Scan(&lastName)
	fmt.Println("Your name is", name, lastName)

}
