package leetcode

func PreorderTraversal(root *TreeNode) []int {
	preorder := []int{}

	if root == nil {
		return preorder
	} else {
		preorder = append(preorder, root.Val)
		preorder = append(preorder, PreorderTraversal(root.Left)...)
		preorder = append(preorder, PreorderTraversal(root.Right)...)
	}

	return preorder
}
