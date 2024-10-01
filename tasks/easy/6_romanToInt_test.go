package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/roman-to-integer/description/

type RomanToInt struct {
	s    string
	want int
}

func TestRomanToInt(t *testing.T) {
	testData := []RomanToInt{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"XXI", 21},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, romanToInt(v.s))
	}
}
