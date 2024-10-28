package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-letter-to-equalize-frequency/description/

type equalFrequencyTask struct {
	str  string
	want bool
}

func TestEqualFrequency(t *testing.T) {
	testData := []equalFrequencyTask{
		{"abcc", true},
		{"aazz", false},
		{"bac", true},
		{"babbdd", false},
		{"aabcbb", false},
		{"zz", true},
		{"cccd", true},
		{"aaaabbbbccc", false},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, equalFrequency(v.str))
	}
}
