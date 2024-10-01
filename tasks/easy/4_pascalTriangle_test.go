package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/pascals-triangle/

type pascalTriangle struct {
	numRows int
	want    [][]int
}

func TestGeneratePascalTriangle(t *testing.T) {
	testData := []pascalTriangle{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, GeneratePascalTriangle(v.numRows))
	}
}
