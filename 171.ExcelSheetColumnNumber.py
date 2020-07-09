# -*- coding:utf8 -*-


# https://leetcode.com/problems/excel-sheet-column-number/

dictA = {}
for i, x in enumerate(list("ABCDEFGHIJKLMNOPQRSTUVWXYZ")):
    dictA[x] = i+1

class Solution:
    def titleToNumber(self, s: str) -> int:
        cnt = 0
        for i in s:
            cnt = cnt * 26 + dictA[i]
        return cnt

if __name__ == "__main__":
    sol = Solution()
    print(sol.titleToNumber("AA"))
