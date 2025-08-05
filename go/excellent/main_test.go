// File: go/excellent/main_test.go

package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("Expected 'even', got '%s'", result)
	}

	result = EvenOrOdd(7)
	if result != "odd" {
		t.Errorf("Expected 'odd', got '%s'", result)
	}
}
