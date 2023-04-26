package leetcode

func IsSymmetric(root *TreeNode) bool {
	return LeftAndRight(root, root)
}
func LeftAndRight(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return LeftAndRight(l.Left, r.Right) && LeftAndRight(l.Right, r.Left)
}
