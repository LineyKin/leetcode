package medium

// https://leetcode.com/problems/reverse-integer/description/

import "math"

func reverse(x int) int {
	hasMinus := false

	if x < 0 {
		hasMinus = true
		x = (-1) * x
	}

	var xSliceRev []int

	for x > 0 {
		lastDigit := x % 10
		xSliceRev = append(xSliceRev, lastDigit)
		x = (x - lastDigit) / 10
	}

	xLen := len(xSliceRev)
	var xRev int
	for i := 0; i < xLen; i++ {
		xRev += xSliceRev[i] * int(math.Pow10(xLen-1-i))
		if hasMinus && xRev > int(math.Pow(2, 31)) {
			return 0
		}

		if !hasMinus && xRev > int(math.Pow(2, 31)-1) {
			return 0
		}
	}

	if hasMinus {
		xRev = (-1) * xRev
	}

	return xRev
}
