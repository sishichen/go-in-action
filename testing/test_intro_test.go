package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

// can also access to other file, access function in test_intro2
func TestCalculateENENR(t *testing.T) {
	if CalculateENEN(2) != 5 {
		t.Error("Expected 2 + 2 to equal 5")
	}
}

type xxx struct {
	input    int
	expected int
}

// test with grouped list
func TestTableCalculate(t *testing.T) {
	var tests = []xxx{
		{2, 4},
		{-1, 1},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
