package easy

import "sort"

// https://leetcode.com/problems/remove-letter-to-equalize-frequency/description/

func equalFrequency(word string) bool {
	myMap := make(map[rune]int8)

	for _, v := range word {
		myMap[v]++
	}

	if len(myMap) == 1 {
		return true
	}

	list := make([]int, len(myMap))
	var i byte

	var max int8
	for _, v := range myMap {
		list[i] = int(v)
		if v > max {
			max = v
		}
		i++
	}

	if max == 1 {
		return true
	}

	sort.Ints(list)

	if list[0] == 1 {
		list := list[1:]
		if list[0] == list[len(list)-1] {
			return true
		}
	}

	for i := 0; i < len(list); i++ {
		if list[i]-int(max) != -1 && i < len(list)-1 {
			return false
		}
	}

	return true
}
