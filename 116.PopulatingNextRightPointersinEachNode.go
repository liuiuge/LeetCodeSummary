package code

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	move(root)
	return root
}

func move(root) {
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}
	root.Left.Next = root.Right
	move(root.Left)
	move(root.Right)
}
