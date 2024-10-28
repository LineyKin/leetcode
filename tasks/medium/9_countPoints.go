package medium

import "math"

// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/

func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	for k, circ := range queries {
		for _, point := range points {
			if calcVectorLen(circ, point) <= float64(circ[2]) {
				res[k]++
			}
		}
	}

	return res
}

func calcVectorLen(circ, point []int) float64 {
	xDiff := float64(circ[0] - point[0])
	yDiff := float64(circ[1] - point[1])

	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}
