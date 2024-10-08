package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/description/?envType=daily-question&envId=2024-10-07

type MinLengthTask struct {
	s    string
	want int
}

func TestMinLength(t *testing.T) {
	testData := []MinLengthTask{
		{"ABFCACDB", 2},
		{"ACBBD", 5},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, MinLength(v.s))
	}
}
