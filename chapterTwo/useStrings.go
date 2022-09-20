package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("EqualFold: %v\n", s.EqualFold("Ricardo", "RIcaRDo"))
	f("EqualFold: %v\n", s.EqualFold("Ricardo", "RIcaRD"))
	f("Index: %v\n", s.Index("Ricardo", "rd"))
	f("Index: %v\n", s.Index("Ricardo", "Rd"))
	f("Prefix: %v\n", s.HasPrefix("Ricardo", "Ri"))
	f("Prefix: %v\n", s.HasPrefix("Ricardo", "ri"))
	f("Suffix: %v\n", s.HasSuffix("Ricardo", "do"))
	f("Suffix: %v\n", s.HasSuffix("Ricardo", "DO"))
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("This is a\tstring!")
	f("Fields: %v\n", len(t))
	f("%s\n", s.Split("abcd ef", ""))
	f("%s\n", s.Replace("abcd ef", "", "_", -1))
	f("%s\n", s.Replace("abcd ef", "", "_", 4))
	f("%s\n", s.Replace("abcd ef", "", "_", 2))
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++")[0])
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TringFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
