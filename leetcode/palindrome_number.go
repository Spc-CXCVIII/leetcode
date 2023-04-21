package leetcode

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	sum := 0
	mul := 1
	tem := x
	for tem > 0 {
		last := tem % 10
		sum *= mul
		sum += last
		tem /= 10
		mul = 10
	}

	return sum == x
}
