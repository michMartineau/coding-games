package main

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i,n := range nums {
		ans[i] = nums[n]
	}
	return ans
}
