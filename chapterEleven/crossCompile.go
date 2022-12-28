package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.GOOS, " ")
	fmt.Println("on a(n)", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}

// env GOOS=linux GOARCH=amd64 go build crossCompile.go
// ./crossCompile
// should fail because it's the wrong os

// env GOOS=darwin GOARCH=amd64 go build crossCompile.go
// ./crossCompile
// should pass
