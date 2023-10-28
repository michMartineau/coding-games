package main

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
