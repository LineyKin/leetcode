package easy

// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/

func differenceOfSums(n int, m int) int {
	var diff int32
	for i := int16(1); i <= int16(n); i++ {
		if i%int16(m) != 0 {
			diff += int32(i)
			continue
		}

		diff -= int32(i)
	}

	return int(diff)
}
