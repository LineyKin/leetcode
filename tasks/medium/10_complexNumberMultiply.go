package medium

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/complex-number-multiplication/

// num1 = a + bi
// num2 = c + di
// num1*num2 = ac - bd + (ad + bc)i

func complexNumberMultiply(num1 string, num2 string) string {
	split1 := strings.Split(num1, "+")
	split2 := strings.Split(num2, "+")

	a, _ := strconv.Atoi(split1[0])
	b, _ := strconv.Atoi(strings.Replace(split1[1], "i", "", 1))
	c, _ := strconv.Atoi(split2[0])
	d, _ := strconv.Atoi(strings.Replace(split2[1], "i", "", 1))

	return fmt.Sprintf("%d+%di", a*c-b*d, a*d+b*c)
}
