package easy

// https://leetcode.com/problems/search-in-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SearchBST(root *TreeNode, val int) *TreeNode {

	if root != nil && root.Val == val {
		return root
	}

	var l, r *TreeNode

	if root.Left != nil {
		l = SearchBST(root.Left, val)
	}

	if root.Right != nil {
		r = SearchBST(root.Right, val)
	}

	if r != nil {
		return r
	}

	if l != nil {
		return l
	}

	return nil
}
