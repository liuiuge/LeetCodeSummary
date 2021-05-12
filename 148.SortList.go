package code

import (
	"sort"
)

// https://leetcode.com/problems/sort-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]int, 0, 50000)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	head = &ListNode{
		Val: arr[0],
	}
	cur := head
	for i := 1; i < len(arr); i++ {
		tmp := &ListNode{
			Val: arr[i],
		}
		cur.Next = tmp
		cur = cur.Next
	}
	return head
}

func sortListSmoothlty(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	pre := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next

	}

	pre.Next = nil

	l1 := sortListSmoothlty(head)
	l2 := sortListSmoothlty(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
