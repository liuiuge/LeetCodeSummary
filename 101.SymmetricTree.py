# -*- coding: utf8 -*-

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def isSymmetric(left, right):
    if not left and not right:
        return True
    if left and right and left.val == right.val:
        return isSymmetric(left.left, right.right) and isSymmetric(left.right, right.left)
    return False

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        return isSymmetric(root.right, root.left)

if __name__ == "__main__":
    pass