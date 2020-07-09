# -*- coding:utf8 -*-

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def helper(array) -> TreeNode:
    if not len(array):
        return None
    mid = len(array) // 2
    root = TreeNode(array[mid])
    if len(array) < 2:
        return root
    root.left = helper(array[0:mid])
    root.right = helper(array[mid+1:len(array)])
    return root


class Solution:

    def sortedArrayToBST(self, nums) -> TreeNode:
        if not len(nums):
            return None
        
        mid = len(nums) // 2
        root = TreeNode(nums[mid])
        if len(nums) < 2:
            return root
        root.left = helper(nums[0:mid])
        root.right = helper(nums[mid+1:len(nums)])
        return root

if __name__ == "__main__":
    a = [-10,-3,0,5,9]
    sol = Solution()
    print(sol.sortedArrayToBST(a).val)
