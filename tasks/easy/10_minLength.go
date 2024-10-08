package easy

import (
	"strings"
)

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/description/?envType=daily-question&envId=2024-10-07

// не оптимальное решение
// переписать с помощью стека

func MinLength2(s string) int {
	stack := []string{}
	for _, byteChar := range s {
		letter := string(byteChar)
		if len(stack) > 0 && (letter == "D" && stack[len(stack)-1] == "C" ||
			letter == "B" && stack[len(stack)-1] == "A") {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, letter)
		}
	}

	return len(stack)
}

func MinLength(s string) int {
	newLen := 0
	for {
		sLen := len(s)
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
		newLen = len(s)

		if sLen == newLen {
			break
		}
	}

	return newLen
}
