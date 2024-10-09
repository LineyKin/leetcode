package easy

// https://leetcode.com/problems/score-of-a-string/

func scoreOfString(s string) int {
	var sum int16
	for i := 1; i < len(s); i++ {
		sum += int16(calcByte(s[i], s[i-1]))
	}

	return int(sum)
}

func calcByte(a, b byte) byte {
	if a > b {
		return a - b
	}

	return b - a
}
