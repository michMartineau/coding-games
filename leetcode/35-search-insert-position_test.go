package main

import "testing"

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target < nums[0] {
		return 0
	}

	for i, v := range nums {
		if target > v {
			continue
		}
		if target <= v {
			return i
		}
	}
	return 0
}

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		input1 []int
		input2 int
		output int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tt := range tests {
		valid := searchInsert(tt.input1, tt.input2)
		if valid != tt.output {
			t.Errorf("for %d and %d not valid %d is expected", tt.input1, tt.input2, tt.output)
		}

	}

}
