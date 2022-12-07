package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

func charByChar(file string) {
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
	set(chars)
}

func lineByLine(file string) {
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
	set(lines)
}

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func wordByWord(file string) {
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
	set(word_count)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1>\n")
		return
	}
	go monitor()
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		lineByLine(os.Args[1])
	}()
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		charByChar(os.Args[1])
	}()
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		wordByWord(os.Args[1])
	}()
	waitGroup.Wait()

	fmt.Println(os.Args[1])
}
