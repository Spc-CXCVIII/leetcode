package leetcode

import "sort"

func SingleNumber(nums []int) int {
	sort.Ints(nums)

	temp := nums[0]
	ans := temp
	for i := 1; i < len(nums); i++ {
		if temp == nums[i] && i <= len(nums)-1 {
			temp = nums[i+1]
			ans = temp
			i++
			continue
		} else {
			ans = temp
			break
		}
	}

	return ans
}
