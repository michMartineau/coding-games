package main

import "testing"

func removeElement(nums []int, val int) int {

	for v := len(nums) - 1; v >= 0; v-- {
		if nums[v] == val {
			nums = append(nums[:v], nums[v+1:]...)
		}
	}
	return len(nums)

}
func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		input1 []int
		input2 int
		output int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tt := range tests {
		valid := removeElement(tt.input1, tt.input2)
		if valid != tt.output {
			t.Errorf("for %d and %d not valid %d is expected", tt.input1, tt.input2, tt.output)
		}

	}

}
