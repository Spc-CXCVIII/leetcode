package leetcode

func SearchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	if len(nums) == 0 {
		return 0
	}

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if low == high {
				return mid + 1
			} else if high-low == 1 {
				low = high
			}
			low = mid
		} else if nums[mid] > target {
			if low == high {
				return mid
			} else if high-low == 1 {
				high = low
			}
			high = mid
		}
	}
	return low
	// [1,3,5,7,9,11,13,15,17,19,21,23,25,27,29]  22
	//                               M             T
}
