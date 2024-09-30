package medium

// https://leetcode.com/problems/reverse-integer/description/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Reverse struct {
	x    int
	want int
}

func TestReverse(t *testing.T) {
	testData := []Reverse{
		{123, 321},
		{-123, -321},
		{120, 21},
		{2147483647, 0},
		{0, 0},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, reverse(v.x))
	}
}
