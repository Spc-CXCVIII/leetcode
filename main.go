package main

import (
	"Test/leetcode"
	"fmt"
)

func main() {
	// list1 := &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 3, Next: &leetcode.ListNode{Val: 3, Next: nil}}}}}
	list2 := &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 2, Next: nil}}}}
	fmt.Println(leetcode.DeleteDuplicates(list2))
}
