package easy

// https://leetcode.com/problems/remove-linked-list-elements/

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		head = removeElements(head.Next, val)
	}

	resList := &ListNode{}
	resList.Val = head.Val
	if head.Next != nil {
		resList.Next = removeElements(head.Next, val)
	}

	return resList
}

// leetcode version

func RemoveElements2(head *ListNode, val int) *ListNode {

	for node, prev := head, head; node != nil; {

		if node.Val == val {
			if node == head {
				head = node.Next
				prev = head
			} else {
				prev.Next = node.Next
			}

		} else {
			prev = node
		}
		node = node.Next
	}
	return head
}
