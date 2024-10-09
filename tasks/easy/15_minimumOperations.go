package easy

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/

func minimumOperations(nums []int) int {
	var minOp byte

	for _, n := range nums {
		if byte(n)%3 != 0 {
			minOp++
		}
	}

	return int(minOp)
}
