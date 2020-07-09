# -*- coding: utf8 -*-

class Solution:
    def generate(self, numRows: int):
        a = [[] for i in range(numRows)]
        for i in range (numRows):
            a[i] = [1 if j == i or j == 0 else a[i-1][j] + a[i-1][j-1] for j in range(i+1)]
        return a

if __name__ == "__main__":
    sol = Solution()
    print(sol.generate(5))