package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/score-of-a-string/

type ScoreOfStringTask struct {
	s    string
	want int
}

func TestScoreOfString(t *testing.T) {
	tetsData := []ScoreOfStringTask{
		{"hello", 13},
		{"zaz", 50},
		{"tgtktpytavhslrnrrxwtbfhqyqronmvlqdxbpsymhgwyb", 374},
	}

	for _, v := range tetsData {
		assert.Equal(t, v.want, scoreOfString(v.s))
	}
}
