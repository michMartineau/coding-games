package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input  int
		output bool
	}{
		{1, false},
		{11, true},
		{121, true},
		{122, false},
	}

	for _, tt := range tests {
		valid := isPalindrome(tt.input)
		if valid != tt.output {
			t.Errorf("for %d not valid %t is expected", tt.input, tt.output)
		}
	}
}
