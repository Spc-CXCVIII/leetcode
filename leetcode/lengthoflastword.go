package leetcode

func LengthOfLastWord(s string) int {
	length := 0
	start := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if !start {
				continue
			} else {
				return length
			}
		} else if s[i] != ' ' {
			start = true
			length++
		}
	}

	return length
}
