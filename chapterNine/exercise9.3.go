package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var count = 0
var min = 10

func random(interval int) int {
	return rand.Intn(interval) + min
}

func handleConnection(c net.Conn, ch chan bool) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		fmt.Println(temp)
		if temp == "STOP" {
			c.Write([]byte("Closing server\n"))
			ch <- true
			break
		}
		interval, err := strconv.Atoi(temp)
		if err == nil {
			c.Write([]byte(strconv.Itoa(random(interval)) + "\n"))
		} else {
			c.Write([]byte("Please introduce a number\n"))
		}
	}
	c.Close()
}

func main() {
	ch := make(chan bool)
	seed := time.Now().Unix()
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(seed)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, ch)
		count++
		for i := range ch {
			fmt.Println(i)
			if i == true {
				close(ch)
				return
			}
		}
	}
}
