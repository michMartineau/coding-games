package main

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	res := make([]int, 0)
	var acc int
	for i := len(digits) - 1; i >= 0; i-- {
		var val int
		if i == len(digits)-1 {
			val++
		}

		val += digits[i] + acc
		acc = 0
		//fmt.Println(i)
		//fmt.Println(val)
		if val == 10 && i != 0 {
			res = append([]int{0}, res...)
			acc = 1
			continue
		}
		if val == 10 && i == 0 {
			res = append([]int{0}, res...)
			res = append([]int{1}, res...)
			continue
		}
		res = append([]int{val}, res...)
	}
	return res
}

// From https://www.youtube.com/watch?v=mrYywKoIq-E&list=PL8-bdB4cHmXzuGvBnFmByq-4yyAy-ceaK&index=5
//func plusOne(digits []int) []int {
//
//	n := len(digits)
//	for i := n - 1; i >= 0; i-- {
//		if digits[i] < 9 {
//			digits[i]++
//			return digits
//		}
//		digits[i] = 0
//
//	}
//	digits = append([]int{1}, digits...)
//	return digits
//}

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{[]int{8}, []int{9}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 8}, []int{9, 9}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{9, 8, 9}, []int{9, 9, 0}},
	}

	for _, tt := range tests {
		valid := plusOne(tt.input)
		if !reflect.DeepEqual(valid, tt.output) {
			t.Errorf("for %d not valid %d is expected", tt.input, tt.output)
		}
	}
}
