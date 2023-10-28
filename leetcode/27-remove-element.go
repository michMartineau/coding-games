package main

func removeElement(nums []int, val int) int {

	for v := len(nums) - 1; v >= 0; v-- {
		if nums[v] == val {
			nums = append(nums[:v], nums[v+1:]...)
		}
	}
	return len(nums)

}
