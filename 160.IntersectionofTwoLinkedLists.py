# -*- coding:utf8 -*-

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if headA is None or headB is None:
            return None
        cntA, cntB = 0, 0
        tmpA, tmpB = headA, headB
        while tmpA is not None:
            cntA += 1
            tmpA = tmpA.next
        while tmpB is not None:
            cntB += 1
            tmpB = tmpB.next
        if cntB > cntA:
            for _ in range(cntB - cntA):
                headB = headB.next
        else:
            for _ in range(cntA - cntB):
                headA = headA.next
        while headA != headB:
            headA = headA.next
            headB = headB.next
        return headA