package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/container-with-most-water/

type ContWithMostWater struct {
	height []int
	want   int
}

func TestMaxArea2(t *testing.T) {
	testData := []ContWithMostWater{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, maxArea2(v.height))
	}
}

func TestMaxArea(t *testing.T) {
	testData := []ContWithMostWater{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, maxArea(v.height))
	}
}
