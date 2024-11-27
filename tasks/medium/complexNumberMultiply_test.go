package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/complex-number-multiplication/

type complexNumberMultiplyTask struct {
	num1 string
	num2 string
	want string
}

func TestComplexNumberMultiply(t *testing.T) {
	testData := []complexNumberMultiplyTask{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, complexNumberMultiply(v.num1, v.num2))
	}
}
