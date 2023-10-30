package main

import "testing"

func removeDuplicates(nums []int) int {

	for v := len(nums) - 1; v >= 0; v-- {
		if v == len(nums)-1 {
			continue
		}
		curr := nums[v]
		prev := nums[v+1]
		if curr == prev {
			nums = append(nums[:v], nums[v+1:]...)
		}

	}
	return len(nums)
}
func TestRemoveDuplicate(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		valid := removeDuplicates(tt.input)
		if valid != tt.output {
			t.Errorf("for %d not valid %d is expected", tt.input, tt.output)
		}

	}

}
