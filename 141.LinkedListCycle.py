# -*- coding:utf8 -*-

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head is None or head.next is None or head.next.next is None:
            return False
        fast = head.next.next
        slow = head.next
        while fast is not None:
            fast = fast.next
            if fast is None:
                return False
            fast = fast.next
            slow = slow.next
            if fast == slow:
                return True
        return False