/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	retlist := make([]int, 0)
	for head != nil {
		retlist = append(retlist, head.Val)
		head = head.Next
	}
	return buildTree(retlist, 0, len(retlist)-1)
}

func buildTree(ret []int, lo, hi int) *TreeNode {
	if hi < lo {
		return nil
	} else if hi == lo {
		return &TreeNode{
			Val: ret[hi],
		}
	}
	mid := (hi + lo) / 2
	root := &TreeNode{
		Val: ret[mid],
	}
	root.Left = buildTree(ret, lo, mid-1)
	root.Right = buildTree(ret, mid+1, hi)
	return root
}