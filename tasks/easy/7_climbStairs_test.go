package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/climbing-stairs/description/

type ClimbStairs struct {
	n    int
	want int
}

func TestClimbStairs(t *testing.T) {
	testData := []ClimbStairs{
		{2, 2},
		{3, 3},
		{44, 1134903170},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, climbStairs(v.n))
	}
}
