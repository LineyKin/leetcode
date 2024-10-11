package medium

// https://leetcode.com/problems/merge-nodes-in-between-zeros/

func MergeNodes(head *ListNode) *ListNode {
	newListNode := &ListNode{}
	tmpNewList := newListNode
	tmp := head
	for tmp.Next != nil {
		if tmp.Val != 0 {
			tmpNewList.Val += tmp.Val
		} else {
			if tmpNewList.Val > 0 {
				tmpNewList.Next = &ListNode{}
				tmpNewList = tmpNewList.Next
			}

		}

		tmp = tmp.Next

	}

	return newListNode
}
