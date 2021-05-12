package code

// https://leetcode.com/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) int {
	cnt := 0
	if root == nil {
		return cnt
	}
	return dfsMinDepth(root, 1)
}

func dfsMinDepth(root *TreeNode, cnt int) int {
	if root.Right == nil && root.Left == nil {
		return cnt
	}
	left, right := 0, 0
	if root.Left != nil {
		left = dfsMinDepth(root.Left, cnt+1)
	}
	if root.Right != nil {
		right = dfsMinDepth(root.Right, cnt+1)
	}
	if right > 0 {
		if right < left {
			return right
		} else if left < 1 {
			return right
		} else {
			return left
		}
	}
	return left
}
