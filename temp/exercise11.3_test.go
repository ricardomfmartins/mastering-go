package main

import (
	"fmt"
	"math"
	"testing"
)

func TestWithItself(t *testing.T) {
	f := fibonacci()
	condition := func(previous_result int) (bool, int) {
		result := f()
		fmt.Println(previous_result, "+", previous_result, "=", result)
		if result == 0 {
			return previous_result == result, result
		}
		if result == 1 && (previous_result == 0 || previous_result == 1) {
			return true, result
		}
		if result == 2 || previous_result == 1 {
			return true, result
		} else {
			temp := math.Round(float64(result)/float64(previous_result)*100) / 100
			fmt.Println(temp, temp >= 1.5, temp <= 1.7, (temp >= 1.5 && temp <= 1.7))
			return (temp >= 1.5 && temp <= 1.7), result
		}
	}
	previous_result := 0
	result := false
	for i := 0; i < 10; i++ {
		result, previous_result = condition(previous_result)
		fmt.Println(result)
		if result == false {
			t.Errorf("Error: %v", result)
		}
	}

}
