# -*- coding:utf8 -*-

# https://leetcode.com/problems/ugly-number/

class Solution:
    def isUgly(self, num: int) -> bool:
        if num == 1:
            return True
        
        while num > 5 and num % 5 == 0:
            num = num // 5
        if num == 5:
            return True
        while num > 3 and num % 3 == 0:
            num = num // 3
        if num == 3:
            return True
        while num > 2 and num % 2 == 0:
            num = num // 2
        
        return num == 2

if __name__ == "__main__":
    sol = Solution()
    print(sol.isUgly(1000))