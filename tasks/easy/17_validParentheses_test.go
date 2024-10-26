package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-parentheses/?envType=problem-list-v2&envId=stack

type validParenthesesTask struct {
	str  string
	want bool
}

func TestIsValid(t *testing.T) {
	testData := []validParenthesesTask{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"]", false},
		{"(", false},
		{"(])", false},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, isValid(v.str))
	}
}
