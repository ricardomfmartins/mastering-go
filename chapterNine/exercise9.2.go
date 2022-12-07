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

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
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
		go handleConnection(c)
		count++
	}
}
