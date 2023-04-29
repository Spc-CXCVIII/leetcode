package leetcode

func TitleToNumber(columnTitle string) int {
	colNum := 0
	mul := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		colNum += (int(columnTitle[i]) - 64) * mul
		mul *= 26
	}

	return colNum
}
