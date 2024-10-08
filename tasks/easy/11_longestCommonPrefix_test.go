package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-common-prefix/description/

type LongestCommonPrefixTask struct {
	words []string
	want  string
}

func TestLongestCommonPrefix(t *testing.T) {
	testData := []LongestCommonPrefixTask{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dog", "dog", "dog"}, "dog"},
		{[]string{"", "", ""}, ""},
		{[]string{""}, ""},
		{[]string{"", "b"}, ""},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, longestCommonPrefix(v.words))
	}
}
