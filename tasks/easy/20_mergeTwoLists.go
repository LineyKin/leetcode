package easy

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	mergeList := &ListNode{}
	if list1.Val <= list2.Val {
		mergeList.Val = list1.Val
		mergeList.Next = MergeTwoLists(list1.Next, list2)
	} else {
		mergeList.Val = list2.Val
		mergeList.Next = MergeTwoLists(list1, list2.Next)
	}

	return mergeList
}
