package code

import (
	"math/rand"
	"time"
)

// https://leetcode.com/problems/linked-list-random-node/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type RNSolution struct {
	length int
	head   *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func RNConstructor(head *ListNode) RNSolution {
	cnt := 0
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	rand.Seed(time.Now().Unix())
	return RNSolution{
		head:   head,
		length: cnt,
	}
}

/** Returns a random node's value. */
func (this *RNSolution) GetRandom() int {
	i, p := rand.Intn(this.length), this.head
	for j := 0; j < i; j++ {
		p = p.Next
	}
	return p.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
