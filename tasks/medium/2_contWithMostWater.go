package medium

// https://leetcode.com/problems/container-with-most-water/

// [1,8,6,2,5,4,8,3,7]
func maxArea(height []int) int {
	var mostWater int
	hLen := len(height)
	for i := 0; i < hLen-1; i++ {
		for j := i + 1; j < hLen; j++ {
			minHeight := height[i]
			if height[i] > height[j] {
				minHeight = height[j]
			}
			waterVal := (j - i) * minHeight
			if waterVal > mostWater {
				mostWater = waterVal
			}
		}
	}

	return mostWater
}

// TODO разобраться и понять, решение не моё
func maxArea2(height []int) int {
	var mostWater int
	i := 0
	j := len(height) - 1
	for i < j {
		minHeight := height[i]
		if height[i] > height[j] {
			minHeight = height[j]
		}
		waterVal := (j - i) * minHeight
		if waterVal > mostWater {
			mostWater = waterVal
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return mostWater
}
