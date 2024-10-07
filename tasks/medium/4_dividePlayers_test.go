package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/?envType=daily-question&envId=2024-10-04

type DividePlayers struct {
	skill []int
	want  int64
}

func TestDividePlayers(t *testing.T) {
	testData := []DividePlayers{
		{[]int{3, 2, 5, 1, 3, 4}, 22},
		{[]int{3, 4}, 12},
		{[]int{1, 1, 2, 3}, -1},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, dividePlayers(v.skill))
	}
}
