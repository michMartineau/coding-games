package main

import (
	"testing"
)

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, val := range operations {
		if val[1] == '+' {
			x++
			continue
		}
		x--
	}
	return x
}

func TestFinalValueAfterOperations(t *testing.T) {
	var tests = []struct {
		input  []string
		output int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}

	for _, tt := range tests {
		valid := finalValueAfterOperations(tt.input)
		if valid != tt.output {
			t.Errorf("for %s not valid %d is expected", tt.input, tt.output)
		}

	}
}
