package leetcode

func RemoveElement(nums []int, val int) int {
	ind_num := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ind_num] = nums[i]
			ind_num++
		}
	}

	return ind_num
}

// [5, 2, 7, 9, 5, 7]  7
// [5, 2, 9]
