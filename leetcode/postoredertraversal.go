package leetcode

func PostorderTraversal(root *TreeNode) []int {
	postorder := []int{}

	if root == nil {
		return postorder
	} else {
		postorder = append(postorder, PostorderTraversal(root.Left)...)
		postorder = append(postorder, PostorderTraversal(root.Right)...)
		postorder = append(postorder, root.Val)
	}

	return postorder
}
