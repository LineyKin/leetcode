package easy

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func DeleteDuplicates2(head *ListNode2) *ListNode2 {
	tmp := head

	if tmp == nil {
		return head
	}

	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return head
}
