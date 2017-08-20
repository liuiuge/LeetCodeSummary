#!/usr/bin/env python
# coding=utf-8
class Solution(object):
    def singleNonDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ans = 0
        for elem in nums:
            ans ^= elem
        return ans
