package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	// Positive test case
	data := []string{
		"abc",
		"beh",
		"chi",
	}
	result := CheckArrayConsistency(&data)
	if result != true {
		t.Errorf("Expected true, but got false")
	}

	// Negative test case - data has different lengths
	data = []string{
		"abc",
		"def",
		"ghij",
	}
	result = CheckArrayConsistency(&data)
	if result != false {
		t.Errorf("Expected false, but got true")
	}

	// Negative test case - data is not symmetric
	data = []string{
		"abc",
		"def",
		"gii",
	}
	result = CheckArrayConsistency(&data)
	if result != false {
		t.Errorf("Expected false, but got true")
	}

	// Positive test case performance
	data = generateString(500)
	start := time.Now()
	result = CheckArrayConsistency(&data)
	end := time.Since(start)
	if result != true {
		t.Errorf("Expected true, but got false")
	}
	fmt.Println(end)
}
