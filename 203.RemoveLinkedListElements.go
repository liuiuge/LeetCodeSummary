package code

// https://leetcode.com/problems/remove-linked-list-elements/

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil {
		pre, cur := head, head
		for cur != nil {
			if cur.Val == val {
				if cur == pre {
					pre = pre.Next
					cur = cur.Next
					head = pre
					continue
				} else {
					cur = cur.Next
					pre.Next = cur
				}
			} else {
				pre = cur
				cur = cur.Next
			}
		}
	}
	return head
}
