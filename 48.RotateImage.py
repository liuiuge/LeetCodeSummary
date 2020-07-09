# -*- coding:utf8 -*-

class Solution:
    def rotate(self, matrix):
        """
        Do not return anything, modify matrix in-place instead.
        """
        for i in range(len(matrix)):
            for j in range(i):
                if i != j:
                    matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        
        for rows in matrix:
            length = len(rows)
            for i in range(length//2):
                rows[i], rows[length-i-1] = rows[length-i-1], rows[i]
        return

if __name__ == "__main__":
    a = [[1,2,3], [4,5,6], [7,8,9]]
    sol = Solution()
    sol.rotate(a)
    print(a)
