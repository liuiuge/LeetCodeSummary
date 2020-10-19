# -*- coding:utf-8 -*-

# https://leetcode.com/problems/matrix-diagonal-sum/

class Solution(object):
    def diagonalSum(self, mat):
        """
        :type mat: List[List[int]]
        :rtype: int
        """
        return sum([mat[i][i]+mat[i][len(mat)-i-1] for i in range(len(mat)) ]) if len(mat)&1 else sum([mat[i][i]+mat[i][len(mat)-i-1] for i in range(len(mat))]) - mat[len(mat)//2][len(mat)//2]


if __name__ == "__main__":
    sol = Solution()
    print(sol.diagonalSum([[1,2,3],
              [4,5,6],
              [7,8,9]]))