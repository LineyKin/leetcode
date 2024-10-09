package easy

// https://leetcode.com/problems/build-array-from-permutation/

func buildArray(nums []int) []int {
	lenNums := int16(len(nums))
	ans := make([]int, lenNums)
	for i := int16(0); i < lenNums; i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}
