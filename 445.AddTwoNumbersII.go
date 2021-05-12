package code

// import (
// 	"fmt"
// )

// https://leetcode.com/problems/add-two-numbers-ii/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var x1, x2 int64 = 0, 0
	for l1 != nil {
		x1 = 10*x1 + int64(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		x2 = 10*x2 + int64(l2.Val)
		l2 = l2.Next
	}
	result := x1 + x2
	p := &ListNode{}
	for result > 0 {
		p.Val = int(result % 10)
		result /= 10
		t := &ListNode{
			Val:  0,
			Next: p,
		}
		p = t
	}
	return p.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	sl, ll := []byte{}, []byte{}
	for l1 != nil {
		sl = append(sl, byte(l1.Val))
		l1 = l1.Next
	}
	for l2 != nil {
		ll = append(ll, byte(l2.Val))
		l2 = l2.Next
	}
	if len(ll) < len(sl) {
		sl, ll = ll, sl
	}
	var add byte = 0
	j := len(ll) - 1
	for i := len(sl) - 1; i >= 0; i-- {
		ret := sl[i] + ll[j] + add
		add = 0
		if ret > 9 {
			ret -= 10
			add = 1
		}
		ll[j] = ret
		j--
	}
	for ; j >= 0 && add > 0; j-- {
		ll[j] = ll[j] + add
		add = 0
		if ll[j] > 9 {
			ll[j] -= 10
			add = 1
		}

	}
	if add > 0 {
		ll = append([]byte{add}, ll...)
	}
	p := &ListNode{}
	for i := len(ll) - 1; i >= 0; i-- {
		p.Val = int(ll[i])
		t := &ListNode{
			Val:  0,
			Next: p,
		}
		p = t
	}
	if p.Next == nil {
		return p
	}
	return p.Next
}
