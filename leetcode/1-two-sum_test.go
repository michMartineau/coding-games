package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input1 []int
		input2 int
		output []int
	}{
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}

	for _, tt := range tests {
		valid := twoSum(tt.input1, tt.input2)
		if !reflect.DeepEqual(valid, tt.output) {
			t.Errorf("for %d not valid %d is expected", tt.input1, tt.output)
		}
	}
}
