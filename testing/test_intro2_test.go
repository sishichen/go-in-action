package main

import (
	"testing"
)

func TestCalculateENEN(t *testing.T) {
	if CalculateENEN(2) != 5 {
		t.Error("Expected 2 + 3 to equal 5")
	}
}

func TestTableCalculateENEN(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 5},
		{-1, 2},
		{9999, 10002},
	}

	for _, test := range tests {
		if output := CalculateENEN(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
