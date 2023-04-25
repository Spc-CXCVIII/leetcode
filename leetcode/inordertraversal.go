package leetcode

func InorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	} else {
		result = append(result, InorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, InorderTraversal(root.Right)...)
	}

	return result
}
