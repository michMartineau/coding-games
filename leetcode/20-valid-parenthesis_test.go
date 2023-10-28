package main

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		chain    string
		expected bool
	}{
		{"", false},
		{"[]", true},
		{"[}]", false},
	}

	for _, tt := range tests {
		valid := isValid(tt.chain)
		if valid != tt.expected {
			t.Errorf("for %s not valid %t is expected", tt.chain, tt.expected)
		}

	}

}
