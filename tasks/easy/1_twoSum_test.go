package easy

// https://leetcode.com/problems/two-sum/description/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TwoSum struct {
	list   []int
	target int
	want   []int
}

func TestTwoSum(t *testing.T) {
	testData := []TwoSum{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{59}, 6, nil},
		{[]int{2, 7, 11, 15}, 900000000000000, nil},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, twoSum(v.list, v.target))
	}
}
