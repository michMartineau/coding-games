package main

import "testing"

// read https://www.code-recipe.com/post/container-with-most-water
func maxArea(height []int) int {
	n := len(height)

	var left, result int
	right := n - 1

	for left < right {
		currArea := min(height[right], height[left]) * (right - left)
		result = max(result, currArea)

		// If the current l height is less than r height,
		// there is no use selecting l for next iteration because, width decreases, so area cant increase
		// So l++(select r, dont select l for next iteration)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, tt := range tests {
		valid := maxArea(tt.input)
		if valid != tt.output {
			t.Errorf("for %d not valid %d is expected, get %d", tt.input, tt.output, valid)
		}
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
