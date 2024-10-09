package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/

type minimumOperationsTask struct {
	nums []int
	want int
}

func TestMinimumOperations(t *testing.T) {
	testData := []minimumOperationsTask{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{3, 6, 9}, 0},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, minimumOperations(v.nums))
	}
}
