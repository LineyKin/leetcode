package easy

// https://leetcode.com/problems/merge-two-binary-trees/

// my version
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil && root2 != nil {
		return root2
	}

	if root1 != nil && root2 == nil {
		return root1
	}

	root1.Val += root2.Val

	if root1.Left != nil {
		root1.Left = MergeTrees(root1.Left, root2.Left)
	}

	if root1.Right != nil {
		root1.Right = MergeTrees(root1.Right, root2.Right)
	}

	if root1.Left == nil && root2.Left != nil {
		root1.Left = root2.Left
	}

	if root1.Right == nil && root2.Right != nil {
		root1.Right = root2.Right
	}

	return root1

}

// leetcode version
func MergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = MergeTrees2(root1.Left, root2.Left)
	root1.Right = MergeTrees2(root2.Right, root1.Right)

	return root1
}
