package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	path := os.Getenv("PATH")
	for i := range arguments {
		file := arguments[i]
		pathSplit := filepath.SplitList(path)
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file? (Avoid Unix non regular files)
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						fmt.Fprintf(os.Stdout, "%s path: %s \n", file, fullPath)
					}
				}
			}
		}
	}
	return
}

// go run exercise2Which.go python3 python
