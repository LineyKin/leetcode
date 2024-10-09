package easy

// https://leetcode.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	lenNums := int16(len(nums))
	ans := make([]int, 2*lenNums)
	for i := int16(0); i < lenNums; i++ {
		ans[i] = nums[i]
		ans[i+lenNums] = nums[i]
	}

	return ans
}
