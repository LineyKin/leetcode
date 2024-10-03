package easy

import (
	"sort"
)

// https://leetcode.com/problems/rank-transform-of-an-array/?envType=daily-question&envId=2024-10-02

func arrayRankTransform(arr []int) []int {
	lenArr := len(arr)
	rankArr := make([]int, lenArr)
	arrSorted := make([]int, lenArr)
	copy(arrSorted, arr)

	// вот не надо изобретать велосипед и пытаться переделать сортировку выбором под нужды задачи
	sort.Ints(arrSorted)

	rankMap := make(map[int]int)
	rank := 1
	for i := 0; i < lenArr; i++ {
		_, exists := rankMap[arrSorted[i]]
		if !exists {
			rankMap[arrSorted[i]] = rank
			rank++
		}
	}

	for i, num := range arr {
		rankArr[i] = rankMap[num]
	}

	return rankArr
}
