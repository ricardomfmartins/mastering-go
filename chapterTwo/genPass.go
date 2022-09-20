package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var MIN int = 0
var MAX int = 94

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("You need to pass one argument with the length of the string")
		return
	}
	length, err := strconv.ParseInt(arguments[1], 10, 64)
	if err != nil {
		fmt.Println("Input needs to be integer")
	}
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	generatedPass := getString(length)
	fmt.Println("Generated password:", generatedPass)
}
