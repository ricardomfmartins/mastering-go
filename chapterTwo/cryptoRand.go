package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
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
	generatedPass, err := generatePass(length)
	if err != nil {
		fmt.Println("Failed to generate ")
	}
	fmt.Println("Generated password:", generatedPass[0:length])
}
