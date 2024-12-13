package easy

// https://leetcode.com/problems/palindrome-linked-list/

// my version
func IsPalindrome(head *ListNode) bool {

	sl := make([]int, 0)

	for head != nil {
		sl = append(sl, head.Val)
		head = head.Next
	}

	for i := 0; i < len(sl)/2; i++ {
		if sl[i] != sl[len(sl)-i] {
			return false
		}
	}

	return true
}
