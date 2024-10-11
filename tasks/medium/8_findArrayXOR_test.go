package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/

type findArrayTask struct {
	nums []int
	want []int
}

func TestFindArray(t *testing.T) {
	testData := []findArrayTask{
		{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
		{[]int{5}, []int{5}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, findArray(v.nums))
	}
}
