package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rank-transform-of-an-array/?envType=daily-question&envId=2024-10-02

type ArrayRankTransform struct {
	list []int
	want []int
}

func TestArrayRankTransform(t *testing.T) {
	testData := []ArrayRankTransform{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
		{[]int{100, 100, 100}, []int{1, 1, 1}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, arrayRankTransform(v.list))
	}
}
