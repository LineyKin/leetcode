package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/build-array-from-permutation/

type buildArrayTask struct {
	arr  []int
	want []int
}

func TestBuildArray(t *testing.T) {
	testData := []buildArrayTask{
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, buildArray(v.arr))
	}
}
