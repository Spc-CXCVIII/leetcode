package leetcode

func RemoveDuplicates(nums []int) int {
	ind_nums := 0 // start index
	cur := nums[0]

	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if cur != nums[i] {
			ind_nums++
			nums[ind_nums] = nums[i]
			cur = nums[i]
		}
	}

	return ind_nums
}

// [2,2,3,7,8,9]
// 2
// dup = 7
