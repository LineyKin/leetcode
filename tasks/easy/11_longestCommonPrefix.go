package easy

// https://leetcode.com/problems/longest-common-prefix/description/

func longestCommonPrefix(strs []string) string {
	totalWords := len(strs)

	if totalWords == 1 {
		return strs[0]
	}

	var commonPrefix string

	if totalWords == 0 {
		return commonPrefix
	}

	byteIndex := 0
	for {
		for wordIndex := 1; wordIndex < totalWords; wordIndex++ {
			if len(strs[wordIndex])-1 < byteIndex || len(strs[wordIndex-1])-1 < byteIndex {
				return commonPrefix
			}

			if strs[wordIndex-1][byteIndex] != strs[wordIndex][byteIndex] {
				return commonPrefix
			}

			if wordIndex == totalWords-1 {
				commonPrefix += string(strs[wordIndex][byteIndex])
			}
		}

		byteIndex++
	}
}
