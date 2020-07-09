# -*- coding: utf8 -*-

class Solution:
    def getRow(self, rowIndex: int):
        # a = [[] for i in range(numRows+1)]
        # for i in range (numRows+1):
        #     a[i] = [1 if j == i or j == 0 else a[i-1][j] + a[i-1][j-1] for j in range(i+1)]
        # return a[numRows]
        prev = [1] + [0] * rowIndex
        curr = [0] * (rowIndex+1)
        for i in range(rowIndex+1):
            for j in range(i+1):
                curr[j] = 1 if j == i or j == 0 else prev[j-1] + prev[j]
            for j in range(i+1):
                prev[j] = curr[j]
        return curr

if __name__ == "__main__":
    sol = Solution()
    print(sol.getRow(5))