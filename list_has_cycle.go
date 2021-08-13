package algos

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	if head == nil || head.Next == nil {
		return false
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
