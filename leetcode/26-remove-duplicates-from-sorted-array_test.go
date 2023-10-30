package main

import "testing"

func removeDuplicates(nums []int) int {
	s := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[s] = nums[i]
			s++
		}
	}
	// fmt.Println(nums)
	// fmt.Println(s)
	// fmt.Println(nums[:s])
	return s
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
