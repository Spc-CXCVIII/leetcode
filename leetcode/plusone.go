package leetcode

func PlusOne(digits []int) []int {
	last := len(digits) - 1
	if digits[last] != 9 {
		digits[last]++
		return digits
	}

	for i := last; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			} else {
				if digits[i-1]+1 == 10 {
					continue
				} else {
					digits[i-1] = digits[i-1] + 1
					break
				}
			}
		}
	}
	return digits
}
