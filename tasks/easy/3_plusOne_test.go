package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/plus-one/description/

type PlusOne struct {
	list []int
	want []int
}

func TestPlusOne(t *testing.T) {
	testData := []PlusOne{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, plusOne(v.list))
	}
}
