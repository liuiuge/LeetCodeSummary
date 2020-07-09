# -*- coding:utf8 -*-

# https://leetcode.com/problems/excel-sheet-column-title/

class Solution:
    def convertToTitle(self, n: int) -> str:
        prepared_list = list("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
        ret = ""
        while n:
            n, idx = divmod(n-1, 26)
            ret += str(prepared_list[idx])
        return ret[::-1]

if __name__ == "__main__":
    sol = Solution()
    ret = sol.convertToTitle(701)
    print(ret)