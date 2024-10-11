package medium

// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/

func InsertGreatestCommonDivisors(head *ListNode) *ListNode {
	tmp := head
	for tmp.Next != nil {
		newNode := &ListNode{}
		newNode.Val = int(GCD(int16(tmp.Val), int16(tmp.Next.Val)))
		newNode.Next = tmp.Next
		tmp.Next = newNode
		tmp = tmp.Next.Next
	}

	return head
}

func GCD(a, b int16) int16 {

	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	c := a % b

	if c == 0 {
		return b
	}

	return GCD(b, c)
}
