package leetcode

import "math"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return 1 + int(math.Max(float64(left), float64(right)))
}
