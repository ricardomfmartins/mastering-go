package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"

	"golang.org/x/sync/semaphore"
)

func charByChar(file string) int {
	f, _ := os.Open(file)
	defer f.Close()

	r := bufio.NewReader(f)
	chars := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		chars += len(line)
	}
	return chars
}

func lineByLine(file string) int {
	f, _ := os.Open(file)
	defer f.Close()

	r := bufio.NewReader(f)
	lines := 0
	for {
		_, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		lines += 1
	}
	return lines
}

// Maximum number of goroutines
var Workers = 2
var sem = semaphore.NewWeighted(int64(Workers))

func wordByWord(file string) int {
	f, _ := os.Open(file)
	defer f.Close()

	r := bufio.NewReader(f)
	word_count := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		re := regexp.MustCompile("[^\\s]+")
		words := re.FindAllString(line, -1)
		word_count += len(words)
	}
	return word_count
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1>\n")
		return
	}
	// Where to store the results
	var results []int

	// Needed by Acquire()
	ctx := context.TODO()

	err := sem.Acquire(ctx, 1)
	if err != nil {
		fmt.Println("Cannot acquire semaphore:", err)
		return
	}
	go func() {
		defer sem.Release(1)
		results = append(results, lineByLine(os.Args[1]))
	}()
	err = sem.Acquire(ctx, 1)
	if err != nil {
		fmt.Println("Cannot acquire semaphore:", err)
		return
	}
	go func() {
		defer sem.Release(1)
		results = append(results, charByChar(os.Args[1]))
	}()
	err = sem.Acquire(ctx, 1)
	if err != nil {
		fmt.Println("Cannot acquire semaphore:", err)
		return
	}
	go func() {
		defer sem.Release(1)
		results = append(results, wordByWord(os.Args[1]))
	}()
	err = sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range results {
		fmt.Print(value, " ")
	}
	fmt.Println(os.Args[1])
}
