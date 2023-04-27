package leetcode

import "fmt"

func ConvertToTitle(columnNumber int) string {
	output := ""

	for columnNumber > 26 {
		var ascii int
		if columnNumber%26 == 0 {
			ascii = 90
			columnNumber -= 1
		} else {
			ascii = (columnNumber % 26) + 64
		}

		char := fmt.Sprintf("%c", ascii)
		output = string(char[0]) + output
		columnNumber /= 26
	}

	ascii := columnNumber + 64
	char := fmt.Sprintf("%c", ascii)
	output = string(char[0]) + output

	return output
}
