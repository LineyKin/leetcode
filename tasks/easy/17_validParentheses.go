package easy

// https://leetcode.com/problems/valid-parentheses/?envType=problem-list-v2&envId=stack

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, v := range s {
		if isOpen(string(v)) {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if isPair(string(v), string(stack[len(stack)-1])) {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) <= 0
}

func isOpen(s string) bool {
	if s == "[" || s == "(" || s == "{" {
		return true
	}

	return false
}

func isPair(s, sStack string) bool {
	if s == "]" && sStack == "[" {
		return true
	}

	if s == ")" && sStack == "(" {
		return true
	}

	if s == "}" && sStack == "{" {
		return true
	}

	return false
}
