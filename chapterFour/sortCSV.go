package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

type F1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type F2 struct {
	Name       string
	Surname    string
	AreaCode   string
	Tel        string
	LastAccess string
}

type Book1 []F1
type Book2 []F2

var d1 = Book1{}
var d2 = Book2{}

func readCSVFile(filePath string) error {
	var firstLine bool = true
	var format1 = true

	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false

			} else {
				return errors.New("Unknown file format!")
			}
			firstLine = false
		}
		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				d1 = append(d1, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					Surname:    line[1],
					AreaCode:   line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				d2 = append(d2, temp)
			}
		}
	}
	return nil
}

func sortData(data interface{}) {
	switch T := data.(type) {
	case Book1:
		d := data.(Book1)
		sort.Sort(Book1(d))
		list(d)
	case Book2:
		d := data.(Book2)
		sort.Sort(Book2(d))
		list(d)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func (a Book1) Len() int {
	return len(a)
}

func (a Book1) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a Book1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Book2) Len() int {
	return len(a)
}

func (a Book2) Less(i, j int) bool {
	if a[i].AreaCode == a[j].AreaCode {
		return a[i].Surname < a[j].Surname
	}
	return a[i].AreaCode < a[j].AreaCode
}

func (a Book2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func list(d interface{}) {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		for _, v := range data {
			fmt.Println(v)
		}
	case Book2:
		data := d.(Book2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func main() {
	var path string
	if len(os.Args) != 2 {
		fmt.Println("Needs a file path")
		return
	} else {
		path = os.Args[1]
	}
	err := readCSVFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(d1) != 0 {
		sortData(d1)
	} else {
		sortData(d2)
	}
}
