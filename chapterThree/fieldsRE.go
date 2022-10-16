package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^\d{9}$`)
	return re.Match(t)
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}
	if !matchNameSur(fields[0]) {
		return false
	}
	if !matchNameSur(fields[1]) {

		return false
	}
	return matchTel(fields[2])
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	s := arguments[1]

	fmt.Println(matchRecord(s))
}
