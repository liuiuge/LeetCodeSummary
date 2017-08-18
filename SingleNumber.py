#!/usr/bin/env python
# coding=utf-8
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cnt = 0
        for num in (nums):
            cnt^=num
        return cnt
