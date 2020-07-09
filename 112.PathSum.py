# -*- coding:utf8 -*-

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def helper1(root, sum, cnt):
    if root is None:
        return False
    cnt += root.val
    if sum > cnt:
        
        ret1 = helper(root.left, sum, cnt)
        ret2 = helper(root.right, sum, cnt)
        print(root.val, ret1, ret2)
        return ret1 | ret2
    elif sum == cnt:
        return True
    else:
        return False

def helper(root, sum, cnt):
    if root is None:
        return False
    cnt += root.val
    if root.left is None and root.right is None and cnt == sum:
        return True
        
    return helper(root.left, sum, cnt) | helper(root.right, sum, cnt)

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if root is None:
            return False
        ret = helper(root, sum, 0)
        return ret

if __name__ == "__main__":
    a = TreeNode(5)
    b = TreeNode(4)
    c = TreeNode(8)
    d = TreeNode(11)
    e = TreeNode(13)
    f = TreeNode(4)
    g = TreeNode(7)
    h = TreeNode(2)
    i = TreeNode(1)
    # a.left = b
    # a.right = c
    # b.left = d
    # c.left = e
    # c.right = f
    # d.left = g
    # d.right = h
    # f.right = i
    i.left = h
    sol = Solution()
    print(sol.hasPathSum(a, 1))
