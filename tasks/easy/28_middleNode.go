package easy

// https://leetcode.com/problems/middle-of-the-linked-list/

// каждые два шага основного списка
// mid будет делать один шаг
// таким образом к концу обхода основного списка
// mid будет по середине, что и будет оьветом
func MiddleNode(head *ListNode) *ListNode {
	mid := head
	counter := 0
	for head != nil {
		counter++
		if counter%2 == 0 {
			mid = mid.Next
		}
		head = head.Next
	}

	return mid
}
