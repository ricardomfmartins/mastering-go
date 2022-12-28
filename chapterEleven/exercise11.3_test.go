package main

import (
	"math"
	"testing"
)

func TestWithItself(t *testing.T) {
	f := fibonacci()
	condition := func(previous_result int) (bool, int) {
		result := f()
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
			return (temp >= 1.5 && temp <= 1.7), result
		}
	}
	previous_result := 0
	result := false
	for i := 0; i < 10; i++ {
		result, previous_result = condition(previous_result)
		if result == false {
			t.Errorf("Error: %v", result)
		}
	}

}
