package easy

// https://leetcode.com/problems/palindrome-number/

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xString := strconv.Itoa(x)
	xRunes := []rune(xString)
	xLen := len(xRunes)

	for i := 0; i < xLen/2; i++ {
		xRunes[i], xRunes[xLen-1-i] = xRunes[xLen-1-i], xRunes[i]
	}

	newX, _ := strconv.Atoi(string(xRunes))

	return newX == x
}
