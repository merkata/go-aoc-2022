package day4

import (
	"strconv"
	"strings"
)

func APFullyContains(input string, option bool) int {
	var result int
	for _, line := range strings.Split(input, "\n") {
		pairs := strings.Split(line, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")
		f1, _ := strconv.Atoi(first[0])
		f2, _ := strconv.Atoi(first[1])
		s1, _ := strconv.Atoi(second[0])
		s2, _ := strconv.Atoi(second[1])
		if !option {
			if (f1 >= s1 && f2 <= s2) || (s1 >= f1 && s2 <= f2) {
				result++
			}
		} else {
			if (f1 >= s1 && f1 <= s2) || (s1 >= f1 && s1 <= f2) || (f2 >= s1 && f2 <= s2) || (s2 >= f1 && s2 <= f2) {
				result++
			}
		}
	}
	return result
}
