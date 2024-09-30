package easy

// https://leetcode.com/problems/plus-one/description/

func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 {

		digits[index]++
		if digits[index] == 10 {
			digits[index] = 0
			index--
		} else {
			break
		}
	}

	if index == -1 {
		return append([]int{1}, digits...)
	}

	return digits
}
