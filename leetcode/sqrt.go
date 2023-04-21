package leetcode

func MySqrt(x int) int {
	temp := 0
	i := 0

	for temp < x {
		temp = i * i
		if temp == x {
			break
		} else if temp > x {
			i--
			break
		}
		i++
	}

	return i
}
