package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// input paramenters are MIN MAX TOTAL SEED
	arguments := os.Args
	min := 0
	max := 99
	total := 100
	seed := time.Now().Unix()

	switch len(arguments) {
	case 3:
		var err, err1 error
		min, err = strconv.Atoi(arguments[1])
		max, err1 = strconv.Atoi(arguments[2])
		if err != nil || err1 != nil {
			fmt.Println("Input needs to be integer")
		}
	case 4:
		var err, err1, err2 error
		min, err = strconv.Atoi(arguments[1])
		max, err1 = strconv.Atoi(arguments[2])
		total, err2 = strconv.Atoi(arguments[3])
		if err != nil || err1 != nil || err2 != nil {
			fmt.Println("Input needs to be integer")
		}
	case 5:
		var err, err1, err2, err3 error
		min, err = strconv.Atoi(arguments[1])
		max, err1 = strconv.Atoi(arguments[2])
		total, err2 = strconv.Atoi(arguments[3])
		seed, err3 = strconv.ParseInt(arguments[4], 10, 64)
		if err != nil || err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Input needs to be integer")
		}
	}

	rand.Seed(seed)
	for i := 0; i < total; i++ {
		fmt.Printf("%d ", random(min, max))
	}
	fmt.Println()
}
