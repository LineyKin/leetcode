package easy

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func Postorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	if root.Children == nil {
		return []int{root.Val}
	}

	list := make([]int, 0, 1)

	for _, childNode := range root.Children {
		list = append(list, Postorder(childNode)...)
	}

	return append(list, root.Val)

}
