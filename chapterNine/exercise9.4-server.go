package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var min = 10
var max = 100

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func handler(c net.Conn) {
	for {
		buf := make([]byte, 128)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Read:", err)
			return
		}

		data := buf[0:n]
		fmt.Println(strings.TrimSpace(string(data)))
		if string(data) == "STOP\n" {
			c.Close()
			os.Exit(0)
		}
		fmt.Print("Server got: ", string(data))
		_, err = c.Write([]byte(strconv.Itoa(random(min, max)) + "\n"))
		if err != nil {
			fmt.Println("Write:", err)
			return
		}
	}
}

func main() {
	seed := time.Now().Unix()
	// Read socket path
	if len(os.Args) == 1 {
		fmt.Println("Need socket path")
		return
	}
	socketPath := os.Args[1]

	// If socketPath exists, delete it
	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing", socketPath)
		err := os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	rand.Seed(seed)

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		go handler(fd)
	}
}
