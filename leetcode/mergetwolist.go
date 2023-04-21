package leetcode

func MergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	mergeNode := &ListNode{}
	cur := mergeNode
	for list1 != nil {
		newNode := &ListNode{}
		if list1.Val < list2.Val {
			newNode.Val = list1.Val
			newNode.Next = nil
			cur.Next = newNode
			cur = cur.Next
			list1 = list1.Next
		} else if list1.Val == list2.Val {
			newNode.Val = list1.Val
			newNode.Next = nil
			cur.Next = newNode
			cur = cur.Next
			list1 = list1.Next

			newNode = &ListNode{}

			newNode.Val = list2.Val
			newNode.Next = nil
			cur.Next = newNode
			cur = cur.Next
			if list2.Next != nil && list1 != nil {
				list2 = list2.Next
			} else if list2.Next == nil && list1 == nil {
				list2 = list2.Next
				break
			} else if list2.Next == nil && list1 != nil {
				cur.Next = list1
				break
			} else if list2.Next != nil && list1 == nil {
				list2 = list2.Next
				cur.Next = list2
				break
			}
		} else {
			newNode.Val = list2.Val
			newNode.Next = nil
			cur.Next = newNode
			cur = cur.Next
			if list2.Next != nil && list1 != nil {
				list2 = list2.Next
			} else if list2.Next == nil && list1 == nil {
				list2 = list2.Next
				break
			} else if list2.Next == nil && list1 != nil {
				cur.Next = list1
				break
			} else if list2.Next != nil && list1 == nil {
				list2 = list2.Next
				cur.Next = list2
				break
			}
		}
	}

	if list1 == nil && list2 != nil {
		cur.Next = list2
	}

	temp := mergeNode.Next
	for temp != nil {
		println(temp.Val)
		temp = temp.Next
	}

	return mergeNode.Next
}
