package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name        string
	Surname     string
	Tel         string
	LastVisited string
}

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int
var CSVFILE = "./csv.file"

var MIN int = 0
var MAX int = 94

func (a PhoneBook) Len() int {
	return len(a)
}

func (a PhoneBook) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastVisited = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func list(reverse bool) {
	if reverse {
		sort.Sort(sort.Reverse(PhoneBook(data)))
	} else {
		sort.Sort(PhoneBook(data))
	}
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^\d{9}$`)
	return re.Match(t)
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}
	if !matchNameSur(fields[0]) {
		return false
	}
	if !matchNameSur(fields[1]) {

		return false
	}
	return matchTel(fields[2])
}

func insert(pS *Entry) error {
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists.", pS.Tel)
	}
	data = append(data, *pS)

	_ = createIndex()
	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func validateArguments(ls []string, n int, example string) bool {
	if len(ls) != n {
		fmt.Println("Usage:", example)
		return false
	}
	return true
}

func initS(N, S, T string) *Entry {
	return &Entry{N, S, T, strconv.FormatInt(time.Now().Unix(), 10)}
}

func readCSVFile(filepath string) error {
	var firstLine bool = true
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	for _, row := range lines {
		if firstLine {
			if len(row) != 4 {
				return errors.New("Unknown file format!")
			}
		}
		temp := Entry{Name: row[0], Surname: row[1], Tel: row[2], LastVisited: row[3]}
		data = append(data, temp)
	}
	err = createIndex()
	if err != nil {
		return err
	}
	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastVisited}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func setCSVFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		CSVFILE = filepath
	}
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", CSVFILE)
	}
	return nil
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}
	err := setCSVFILE()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Differentiate between the commands
	switch arguments[1] {
	// The search command
	case "search":
		if !validateArguments(arguments, 3, "search Surname") {
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	// The list command
	case "list":
		if !validateArguments(arguments, 2, "list") {
			return
		}
		list(false)
	// The reverse command
	case "reverse":
		if !validateArguments(arguments, 2, "reverse") {
			return
		}
		reverse := false
		if arguments[1] == "reverse" {
			reverse = true
		}
		list(reverse)
	// Anything that is not a match
	case "insert":
		if !validateArguments(arguments, 3, "insert recordData") {
			return
		}
		if !matchRecord(arguments[2]) {
			fmt.Println("Record with wrong format should be Name,Surname,PhoneNumber.")
			return
		}
		fields := strings.Split(arguments[2], ",")
		record := initS(fields[0], fields[1], fields[2])
		err := insert(record)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "delete":
		if !validateArguments(arguments, 3, "delete surname") {
			return
		}
		err := deleteEntry(arguments[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		list(false)
	default:
		fmt.Println("Not a valid option")
	}

}
