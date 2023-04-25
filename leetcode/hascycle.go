package leetcode

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil { // fast = nil จะเกิดเฉพาะ link list แบบไม่ใช่ cycle
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
