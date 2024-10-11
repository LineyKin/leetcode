package medium

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/

func findArray(pref []int) []int {
	maxIndex := int32(len(pref)) - 1
	for i := maxIndex; i > 0; i-- {
		pref[i] ^= pref[i-1]
	}

	return pref
}
