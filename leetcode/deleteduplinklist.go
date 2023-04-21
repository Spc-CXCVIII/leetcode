package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	new_node := &ListNode{Val: head.Val, Next: nil}
	new_node_head := new_node

	for temp != nil {
		if temp.Next == nil {
			break
		} else if temp.Next != nil && temp.Next.Val != new_node.Val {
			insert_node := &ListNode{Val: temp.Next.Val, Next: nil}
			new_node.Next = insert_node
			new_node = new_node.Next
			temp = temp.Next
		} else if temp.Next != nil && temp.Next.Val == new_node.Val {
			temp = temp.Next
		}
	}

	return new_node_head
}
