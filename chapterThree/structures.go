package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

// Initialized by Go compiler
func goStruct() Entry {
	return Entry{}
}

// Initialized by the user
func userStruct(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// Initialized by Go return pointer
func goPointerStruct() *Entry {
	t := &Entry{}
	return t
}

// Initialized by user - returns pointer
func userPointerStruct(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "John Doe", Year: 2000}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	gs1 := goStruct()
	gps1 := goPointerStruct()
	fmt.Println("gs1", gs1, "gps1", *gps1)
	us2 := userStruct("Darth", "Vader", 4053)
	ups2 := userPointerStruct("Darth", "Vader", 4053)
	fmt.Println("us2", us2, "ups2", *ups2)
	fmt.Println("Year:", gs1.Year, us2.Year, gps1.Year, ups2.Year)

	ngps3 := new(Entry)
	fmt.Println("ngos3", ngps3)

}
