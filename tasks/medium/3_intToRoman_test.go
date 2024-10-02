package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/integer-to-roman/description/

type IntToRoman struct {
	num  int
	want string
}

func TestIntToRoman(t *testing.T) {
	testData := []IntToRoman{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{21, "XXI"},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, intToRoman(v.num))
	}
}
