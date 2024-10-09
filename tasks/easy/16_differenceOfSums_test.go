package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/

type differenceOfSumsTask struct {
	n, m, want int
}

func TestDifferenceOfSums(t *testing.T) {
	testData := []differenceOfSumsTask{
		{10, 3, 19},
		{5, 6, 15},
		{776, 9, 234138},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, differenceOfSums(v.n, v.m))
	}
}
