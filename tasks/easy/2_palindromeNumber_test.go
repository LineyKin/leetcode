package easy

// https://leetcode.com/problems/palindrome-number/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PalindromeNumber struct {
	x    int
	want bool
}

func TestPalindromeNumber(t *testing.T) {
	testData := []PalindromeNumber{
		{121, true},
		{12561, false},
		{-121, false},
		{0, true},
		{1991, true},
		{19791, true},
		{1234598657, false},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, isPalindrome(v.x))
	}
}
