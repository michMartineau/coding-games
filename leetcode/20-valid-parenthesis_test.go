package main

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"", false},
		{"[]", true},
		{"[}]", false},
	}

	for _, tt := range tests {
		valid := isValid(tt.input)
		if valid != tt.output {
			t.Errorf("for %s not valid %t is expected", tt.input, tt.output)
		}

	}

}
