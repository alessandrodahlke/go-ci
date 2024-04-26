package main

import (
	"testing"
)


func TestSum(t *testing.T) {
	// Test case 1
	result := sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}

	// Test case 2
	result = sum(-5, 10)
	expected = 5
	if result != expected {
		t.Errorf("Sum(-5, 10) = %d; expected %d", result, expected)
	}

}