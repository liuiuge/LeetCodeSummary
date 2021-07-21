package code

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	postON(root, &result)
	return result
}

func postON(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for i := range root.Children {
		postON(root.Children[i], result)
	}
	*result = append(*result, root.Val)
}
