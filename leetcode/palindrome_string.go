package leetcode

import (
	"regexp"
	"strings"
)

func IsPalindromeString(s string) bool {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := regex.ReplaceAllString(s, "")

	for i := 0; i < len(result)/2; i++ {
		if !strings.EqualFold(string(result[i]), string(result[len(result)-i-1])) {
			return false
		}
	}

	return true
}
