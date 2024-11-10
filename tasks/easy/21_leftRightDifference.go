package easy

// https://leetcode.com/problems/left-and-right-sum-differences/description/

func leftRightDifference(nums []int) []int {
	ans := make([]int, len(nums))
	var totalSum int
	for _, v := range nums {
		totalSum += v
	}

	var r, l int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			r = totalSum - nums[i]
			l = 0
			ans[0] = r
			continue
		} else {
			r -= nums[i]
			l += nums[i-1]
		}

		if r > l {
			ans[i] = r - l
		} else {
			ans[i] = l - r
		}
	}

	return ans
}
