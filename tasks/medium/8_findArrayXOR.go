package medium

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/

func findArray(pref []int) []int {
	lp := len(pref)
	if lp == 1 {
		return pref
	}

	for i := 1; i < lp; i++ {
		for j := 0; j < i; j++ {
			pref[i] ^= pref[j]
		}
	}

	return pref
}
