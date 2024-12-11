package easy

// https://leetcode.com/problems/evaluate-boolean-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func EvaluateTree(root *TreeNode) bool {
	switch root.Val {
	case 3:
		return EvaluateTree(root.Left) && EvaluateTree(root.Right)
	case 2:
		return EvaluateTree(root.Left) || EvaluateTree(root.Right)
	case 1:
		return true
	default:
		return false
	}
}
