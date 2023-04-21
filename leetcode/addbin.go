package leetcode

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	for len(a) != len(b) {
		if len(a) < len(b) {
			a = "0" + a
		} else if len(b) < len(a) {
			b = "0" + b
		}
	}

	ans := ""
	ove := 0
	// 000010100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101
	// 110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011
	//                                                                                      10011000000000

	for i := len(a) - 1; i >= 0; i-- {
		aInt, _ := strconv.Atoi(string(a[i]))
		bInt, _ := strconv.Atoi(string(b[i]))
		if aInt+bInt+ove == 2 {
			ans = "0" + ans
			if i == 0 {
				ans = "1" + ans
				break
			} else {
				apInt, _ := strconv.Atoi(string(a[i-1]))
				bpInt, _ := strconv.Atoi(string(b[i-1]))
				if apInt+bpInt+1 >= 2 {
					ove = 1
				} else {
					ove = 0
					ans = "1" + ans
					i--
				}
			}
		} else if aInt+bInt+ove == 3 {
			ans = "1" + ans
			if i == 0 {
				ans = "1" + ans
				break
			} else {
				apInt, _ := strconv.Atoi(string(a[i-1]))
				bpInt, _ := strconv.Atoi(string(b[i-1]))
				if apInt+bpInt+1 >= 2 {
					ove = 1
				} else {
					ove = 0
					ans = "1" + ans
					i--
				}
			}
		} else if aInt+bInt+ove == 1 {
			ans = "1" + ans
		} else {
			ans = "0" + ans
		}
	}

	return ans
}
