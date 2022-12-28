package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

var BUFFERSIZE int64

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func copyIoUtil(src, dst string) (int64, error) {

	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		fmt.Println("Error creating", dst)
		fmt.Println(err)
		return 0, err
	}
	return int64(len(input)), err
}

func copyBuf(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return 0, fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}
	buf := make([]byte, 200)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return 0, err
		}
	}
	return 0, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Please provide three command line arguments!")
		return
	}

	sourceFile := os.Args[1]
	destinationFile := os.Args[2]
	cmd, err := strconv.Atoi(os.Args[3])

	var nBytes int64

	if err == nil {
		switch cmd {
		case 1:
			nBytes, err = copy(sourceFile, destinationFile)
		case 2:
			nBytes, err = copyIoUtil(sourceFile, destinationFile)
		case 3:
			nBytes, err = copyBuf(sourceFile, destinationFile)
		default:
			fmt.Printf("Not supported copy command index: %T\n", cmd)
		}
	}

	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}
}
