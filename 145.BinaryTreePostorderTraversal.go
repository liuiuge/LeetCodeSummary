package code

// https://leetcode.com/problems/binary-tree-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postO(root, result)
	return result
}

func postO(root *TreeNode, result []int) {
	if root == nil {
		return
	}
	postO(root.Left, result)
	postO(root.Right, result)
	result = append(result, root.Val)
}
