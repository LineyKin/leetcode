package easy

// https://leetcode.com/problems/reverse-linked-list/

// my version
func ReverseList(head *ListNode) *ListNode {
	s := make([]int, 0)

	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	return buildRevList(s, len(s)-1)
}

func buildRevList(s []int, index int) *ListNode {
	if index < 0 {
		return nil
	}

	var revList ListNode
	revList.Val = s[index]
	revList.Next = buildRevList(s, index-1)

	return &revList
}

// leetcode version
func ReverseList2(head *ListNode) *ListNode {

	// original list
	checkPoint := head

	// reverse list
	var prev *ListNode = nil

	for checkPoint != nil {
		next := checkPoint.Next
		checkPoint.Next = prev
		prev = checkPoint
		checkPoint = next
	}
	return prev
}
