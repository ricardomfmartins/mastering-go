package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
)

var ERR error
var SOURCEFILE string = "1000.txt"

func benchmarkMove(b *testing.B, src string, command int) {
	filename := path.Join(os.TempDir(), src)
	filename = filename + "-" + strconv.Itoa(command)
	var err error
	if err == nil {
		switch command {
		case 1:
			_, err = copy(src, filename)
		case 2:
			_, err = copyIoUtil(src, filename)
		case 3:
			_, err = copyBuf(src, filename)
		default:
			fmt.Printf("Not supported copy command index: %T\n", command)
		}
	}
	ERR = err

	err = os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
	ERR = err
}

func BenchmarkMove1(b *testing.B) {
	benchmarkMove(b, SOURCEFILE, 1)
}

func BenchmarkMove2(b *testing.B) {
	benchmarkMove(b, SOURCEFILE, 2)
}

func BenchmarkMove3(b *testing.B) {
	benchmarkMove(b, SOURCEFILE, 3)
}
