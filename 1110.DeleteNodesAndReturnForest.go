package code

// https://leetcode.com/problems/delete-nodes-and-return-forest/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	result := make([]*TreeNode, 0, 1)
	if root == nil {
		return result
	}
	var check_bit struct{}
	check_map := make(map[int]struct{}, len(to_delete))
	for i := range to_delete {
		check_map[to_delete[i]] = check_bit
	}
	splitTree(check_map, root, &result, true)
	return result
}

func splitTree(check_map map[int]struct{}, root *TreeNode, result *[]*TreeNode, isNil bool) bool {
	if root == nil {
		return
	}
	if _, ok := check_map[root.Val]; ok {
		if root.Left != nil {
			keep := splitTree(check_map, root.Left, result, true)
			if !keep {
				root.Left = nil
			}
		}
		if root.Right != nil {
			keep := splitTree(check_map, root.Right, result, true)
			if !keep {
				root.Right = nil
			}
		}
		root = nil
		return false
	}

	keep := splitTree(check_map, root.Left, result, false)
	if !keep {
		root.Left = nil
	}
	keep = splitTree(check_map, root.Right, result, false)
	if !keep {
		root.Right = nil
	}
	if isNew {
		*result = append(*result, root)
	}
	return true
}
