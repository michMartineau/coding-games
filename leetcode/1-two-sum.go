package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i, v := range nums {

		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				res[0] = i
				res[1] = j
				break
			}
		}
	}
	return res
}
