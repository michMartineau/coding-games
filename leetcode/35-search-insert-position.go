package main

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
