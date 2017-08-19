#!/usr/bin/env python
# coding=utf-8
class Solution:
    def findTheDifference(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        ans = 0
        for elem in s + t:
            ans ^= ord(elem)
        return chr(ans)
