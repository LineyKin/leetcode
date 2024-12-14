package easy

// https://leetcode.com/problems/linked-list-cycle/

func HasCycle(head *ListNode) bool {
	myMap := make(map[*ListNode]bool)

	for head != nil {
		_, ok := myMap[head.Next]
		if ok {
			return true
		}

		myMap[head.Next] = true
		head = head.Next
	}

	return false
}
