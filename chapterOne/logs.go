package main

import (
	"log"
	"os"
)

func mai	n() {
	log.Println("Everything is good")
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
