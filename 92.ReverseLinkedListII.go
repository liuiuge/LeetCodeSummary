package code

// https://leetcode.com/problems/reverse-linked-list-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var start, pre, cur, pro, tail *ListNode = nil, nil, nil, nil, nil
	cur = head
	for i := 1; i < m; i++ {
		start = cur
		cur = cur.Next
	}
	tail = cur
	for i := m; i <= n; i++ {
		pro = cur.Next
		cur.Next = pre
		pre = cur
		cur = pro
	}
	if start == nil {
		if cur == nil {
			return pre
		} else {
			head.Next = pro
			return pre
		}
	} else {
		start.Next = pre
		if cur != nil {
			tail.Next = pro
		}
	}
	return head
}
