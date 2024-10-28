package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/distribute-money-to-maximum-children/

type distMoneyTask struct {
	money    int
	children int
	want     int
}

func TestDistMoney(t *testing.T) {
	testData := []distMoneyTask{
		{20, 3, 1},
		{16, 2, 2},
		{1, 2, -1},
		{2, 2, 0},
		{3, 2, 0},
		{16, 10, 0},
		{17, 2, 1},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, distMoney(v.money, v.children))
	}
}
