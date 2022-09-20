package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}
	dateString := os.Args[1]
	now, _ := time.Parse("_2 Jan 2006 15:04 MST", dateString)
	loc, _ := time.LoadLocation("Europe/Lisbon")
	fmt.Printf("Lisbon time: %s\n", now.In(loc))
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York time: %s\n", now.In(loc))
}

//go run convertTimes.go "14 Dec 2020 19:20 EET"
