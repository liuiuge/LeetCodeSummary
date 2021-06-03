package code

// https://leetcode.com/problems/balanced-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, ok := level(root)
	return ok
}

func level(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	llvl, lok := level(root.Left)
	rlvl, rok := level(root.Right)
	if !(lok && rok) {
		return 0, false
	}
	ret := llvl - rlvl
	if ret*ret < 2 {
		return bigger(llvl, rlvl) + 1, true
	}
	return 0, false
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
