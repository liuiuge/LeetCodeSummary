package code

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}
	tmp, rev := []*TreeNode{root}, false
	for len(tmp) > 0 {
		vtmp, ntmp := make([]int, 0, len(tmp)), make([]*TreeNode, 0, 2*len(tmp))
		for i := range tmp {
			idx := i
			if rev {
				idx = len(tmp) - 1 - i
			}
			vtmp = append(vtmp, tmp[idx].Val)
			if tmp[i].Left != nil {
				ntmp = append(ntmp, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				ntmp = append(ntmp, tmp[i].Right)
			}
		}
		rev = !rev
		result = append(result, vtmp)
		tmp = ntmp
	}
	return result
}
