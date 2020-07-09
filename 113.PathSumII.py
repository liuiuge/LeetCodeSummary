# -*- coding:utf8 -*-
import copy
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def helper(root, result, sum1, current):
    if root is None:
        return
    print(current, result)
    current.append(root.val)
    if root.left is None and root.right is None and sum(current) == sum1:
        result.append(current)
        return
    cur_l = copy.deepcopy(current)
    helper(root.left, result, sum1, cur_l)
    cur_r = copy.deepcopy(current)
    helper(root.right, result, sum1, cur_r)



class Solution:
    def pathSum(self, root: TreeNode, sum1: int):
        def dfs(prev, root, tgt):
            print(prev, result)
            if not root:
                return
            else:
                if not root.left and not root.right and root.val == tgt:
                    prev.append(root.val)
                    result.append(list(prev))
                else:
                    prev.append(root.val)
                    dfs(prev, root.left, tgt - root.val)
                    dfs(prev, root.right, tgt - root.val)
                prev.pop()

        result = []
        # helper(root, result, sum1, [])
        dfs([], root, sum1)
        return result

if __name__ == "__main__":
    a = TreeNode(5)
    b = TreeNode(4)
    c = TreeNode(8)
    d = TreeNode(11)
    e = TreeNode(13)
    f = TreeNode(4)
    g = TreeNode(2)
    h = TreeNode(5)
    i = TreeNode(1)
    a.left = b
    a.right = c
    b.left = d
    c.left = e
    c.right = f
    d.left = g
    f.left = h
    f.right = i
    sol = Solution()
    print(sol.pathSum(a, 22))