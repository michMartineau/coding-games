package main

func numIdenticalPairs(nums []int) int {
	var count int
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v == nums[j] {
				count++
			}
		}
	}
	return count
}
