package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/concatenation-of-array/

type getConcatenationTask struct {
	arr  []int
	want []int
}

func TestGetConcatenation(t *testing.T) {
	testData := []getConcatenationTask{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, getConcatenation(v.arr))
	}
}
