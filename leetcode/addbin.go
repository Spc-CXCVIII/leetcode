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
