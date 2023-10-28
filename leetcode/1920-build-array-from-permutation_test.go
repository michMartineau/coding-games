package main

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}

	for _, tt := range tests {
		valid := buildArray(tt.input)
		if !reflect.DeepEqual(valid, tt.output) {
			t.Errorf("for %d not valid %d is expected", tt.input, tt.output)
		}
	}
}
