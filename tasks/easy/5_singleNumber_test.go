package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/single-number/description/

type SingleNumber struct {
	list []int
	want int
}

func TestSingleNumber(t *testing.T) {
	testData := []SingleNumber{
		{[]int{1}, 1},
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1, 4, 2, 1, 4}, 2},
		{[]int{4, 1, 2, 1, 2, 4, 9}, 9},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, singleNumber(v.list))
	}
}

func TestSingleNumberXOR(t *testing.T) {
	testData := []SingleNumber{
		{[]int{1}, 1},
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1, 4, 2, 1, 4}, 2},
		{[]int{4, 1, 2, 1, 2, 4, 9}, 9},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, singleNumberXOR(v.list))
	}
}
