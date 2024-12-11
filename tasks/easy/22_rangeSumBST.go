package easy

// https://leetcode.com/problems/range-sum-of-bst/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return 0
	}

	var sum int

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	sum += RangeSumBST(root.Left, low, high)

	sum += RangeSumBST(root.Right, low, high)

	return sum

}
