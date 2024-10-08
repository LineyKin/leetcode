package medium

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var inMind int
	var endDigit int
	list := &ListNode{}
	tmp := list
	for isListsNotFinished(l1, l2) || endDigit != 0 {
		sum := getVal(l1) + getVal(l2) + inMind

		if endDigit != 0 {
			sum = endDigit
			endDigit = 0
		}

		if sum > 9 {
			if !isLast(l1, l2) {
				sum -= 10
				inMind = 1
			} else {
				endDigit = sum / 10
				sum = sum % 10
			}
		} else {
			inMind = 0
		}

		tmp.Val = sum

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if isListsNotFinished(l1, l2) || endDigit != 0 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}

	return list
}

// хотя бы один список ещё не пройден до конца`]
func isListsNotFinished(l1, l2 *ListNode) bool {
	return l1 != nil || l2 != nil
}

func isLast(l1, l2 *ListNode) bool {

	// l1 пуст l2 последний элемент
	if l1 == nil {
		return l2 != nil && l2.Next == nil
	}

	// l2 пуст l1 последний элемент
	if l2 == nil {
		return l1 != nil && l1.Next == nil
	}

	// оба списка на последнем элементе
	return l2.Next == nil && l1.Next == nil
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}

	return l.Val
}
