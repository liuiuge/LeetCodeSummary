# -*- coding:utf8 -*-
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def swapPairs(self, head):
        if head is None or head.next is None:
            return head
        
        newhead = head.next
        head.next = newhead.next
        newhead.next = head
        tmp = head.next
        pre = head
        while tmp is not None and tmp.next is not None:
            print(tmp.val, pre.val)
            tmpnext = tmp.next
            tmp.next = tmpnext.next
            tmpnext.next = tmp
            pre.next = tmpnext
            pre = tmp
            tmp = tmp.next
        return newhead

if __name__ == "__main__":
    a, b, c, d = ListNode(1),ListNode(2),ListNode(3),ListNode(4)
    a.next= b
    b.next= c
    c.next= d
    sol = Solution()
    x = sol.swapPairs(a)
    while x is not None:
        print(x.val)
        x = x.next