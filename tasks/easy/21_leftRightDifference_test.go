package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/left-and-right-sum-differences/description/

type leftRightDifferenceTask struct {
	nums []int
	want []int
}

func TestLeftRightDifference(t *testing.T) {
	testData := []leftRightDifferenceTask{
		{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
		{[]int{1}, []int{0}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, leftRightDifference(v.nums))
	}
}
