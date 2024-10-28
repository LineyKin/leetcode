package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/

type countPointsTask struct {
	points  [][]int
	queries [][]int

	want []int
}

func TestCountPoints(t *testing.T) {
	testData := []countPointsTask{
		{
			[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			[][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			[]int{3, 2, 2},
		},

		{
			[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			[][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
			[]int{2, 3, 2, 4},
		},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, countPoints(v.points, v.queries))
	}
}
