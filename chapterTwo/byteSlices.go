package main

import "fmt"

func main() {
	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)
	fmt.Printf("Byte slice: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
	fmt.Println("Length of b:", len(b))
}
