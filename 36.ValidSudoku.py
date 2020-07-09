# -*- coding: utf8 -*-

'''
    判断数独是否有效
'''


class Solution:
    def isValidSudoku(self, board) -> bool:
        if len(board) != 9 or len(board[0]) != 9:
            return False
        for rows in board:
            for i in rows:
                if rows.count(i) > 1 and i != ".":
                    return False
        
        for j in range(9):
            list1 = [i[j] for i in board]
            for x in list1:
                if list1.count(x) > 1 and x != ".":
                    return False
        for i in [0, 3, 6]:
            for j in [0, 3, 6]:
                list1 = [board[x][y] for x in range(i, i+3) for y in range(j, j+3)]
                for x in list1:
                    if list1.count(x) > 1 and x != ".":
                        return False
        return True


if __name__ == "__main__":
    board = [
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
    sol = Solution()
    print(sol.isValidSudoku(board))