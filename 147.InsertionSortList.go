package code

// https://leetcode.com/problems/insertion-sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pcur, plast := head.Next, head
	for pcur != nil {
		if plast.Val > pcur.Val {
			plast.Next = pcur.Next
			if pcur.Val < head.Val {
				pcur.Next = head
				head = pcur
			} else {
				tmp := head
				for {
					if tmp.Next.Val > pcur.Val {
						break
					}
					tmp = tmp.Next
				}
				pcur.Next = tmp.Next
				tmp.Next = pcur
			}
			pcur = plast.Next
		} else {
			plast = pcur
			pcur = pcur.Next
		}
	}
	return head
}
