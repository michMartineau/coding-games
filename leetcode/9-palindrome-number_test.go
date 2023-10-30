package main

import (
	"testing"
)

func reversedNum(x int) int {
	var reversedNum int
	for x > 0 {
		// add the last digit, next loop multiply it by 10
		reversedNum = reversedNum*10 + x%10
		// keep the number without the last digit for the next loop
		x = x / 10
		//fmt.Println(reversedNum)
		//fmt.Println(tmp)
	}
	return reversedNum
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}
	return reversedNum(x) == x
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input  int
		output bool
	}{
		{1, true},
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
