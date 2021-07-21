# -*- coding:utf8 -*-

# https://leetcode.com/problems/compare-version-numbers/

class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        v1, v2 = [int(n) for n in version1.split('.')], [
            int(n) for n in version2.split('.')]
        i = 0
        while i < len(v1) and i < len(v2):
            if v1[i] > v2[i]:
                return 1
            elif v1[i] < v2[i]:
                return -1
            i += 1
        if len(v1) > i and sum(v1[i:]) > 0:
            return 1
        if len(v2) > i and sum(v2[i:]) > 0:
            return -1
        return 0
        
