package easy

// https://leetcode.com/problems/single-number/description/

// [4,1,2,1,2]

func singleNumber(nums []int) int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i] == nums[j] {
			nums = (append(nums[i+1:j], nums[j+1:]...))
			j = len(nums) - 1
		} else {
			j--
		}
	}

	return nums[i]
}

// TODO решение не моё, разобраться`
func singleNumberXOR(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
