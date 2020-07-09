# -*- coding:utf8 -*- 

'''
    旋转链表
'''

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution1:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if head is None or head.next is None:
            return head
        length = self.getLength(head)

        for k in range (k/length):
            head = self.helper(head)
        return head
    
    def getLength(self, head):
        tmp = head
        cnt = 0
        while tmp:
            cnt += 1
            tmp = tmp.next
        return cnt
    
    def helper(self, head):
        tmp = head
        tail = tmp.next
        while True:
            if tail.next is not None:
                # print(tail.val)
                tail = tail.next
                tmp = tmp.next
            else:
                break
        tail.next = head
        tmp.next = None
        # printList(tail)
        return tail

class Solution2:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if head is None or head.next is None:
            return head
        length = self.getLength(head)
        num = length - k % length
        tmp = head
        pre = ListNode()
        n = 0
        while n < num:
            pre = tmp
            tmp = tmp.next
            n += 1
        pre.next = None
        end = tmp
        while end.next:
            end = end.next
        end.next = head
        return tmp 

    def getLength(self, head):
        tmp = head
        cnt = 0
        while tmp:
            cnt += 1
            tmp = tmp.next
        return cnt

def printList(head):
    tmp = head
    while tmp:
        print(tmp.val)
        tmp = tmp.next

if __name__ == "__main__":
    L1 = [1,2,3,4,5]
    head = ListNode(L1[0])
    # printList(head)
    tmp = head
    for i in range(1, len(L1)):
        tmp.next = ListNode(L1[i])
        tmp = tmp.next
    sol = Solution2()
    head = sol.rotateRight(head, 1)
    printList(head)