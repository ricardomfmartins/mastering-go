package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

var MIN int = 0
var MAX int = 94

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

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

func populate(n int, s []Entry) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
	}

}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	populate(100, data)

	// Differentiate between the commands
	switch arguments[1] {
	// The search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	// The list command
	case "list":
		list()
	// Anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
