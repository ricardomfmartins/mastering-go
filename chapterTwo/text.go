package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First character:", string(aString[0]))

	// rune needs single quotes
	r := '€'
	fmt.Println("As an int32 value", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// prinet as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// print as characters
	for _, v := range aString {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
