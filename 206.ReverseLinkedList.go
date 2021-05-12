package code

// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	c_node, p_node, o_node := &ListNode{}, &ListNode{}, &ListNode{}
	c_node, p_node, o_node = nil, head, head.Next
	for p_node.Next != nil {
		o_node = p_node.Next
		p_node.Next = c_node
		c_node = p_node
		p_node = o_node
	}
	p_node.Next = c_node
	return p_node
}
