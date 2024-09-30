package easy

import "math"

// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	len := len(nums)
	if len < 2 || len > 10000 {
		return nil
	}

	if math.Abs(float64(target)) > math.Pow10(9) {
		return nil
	}

	answer := make([]int, 0, 2)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				answer = []int{i, j}
				return answer
			}
		}
	}

	return nil
}
