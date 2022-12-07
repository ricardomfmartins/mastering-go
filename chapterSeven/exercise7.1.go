package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

func charByChar(file string, c chan int) {
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
	c <- chars
}

func lineByLine(file string, c chan int) {
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
	c <- lines
}

func wordByWord(file string, c chan int) {
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
	c <- word_count
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1>\n")
		return
	}
	channels := make(chan int, 3)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		lineByLine(os.Args[1], channels)
	}(channels)
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		charByChar(os.Args[1], channels)
	}(channels)
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		wordByWord(os.Args[1], channels)
	}(channels)
	waitGroup.Wait()
	for {
		select {
		case channel := <-channels:
			fmt.Print(channel, " ")
		default:
			fmt.Println(os.Args[1])
			return
		}
	}
}
