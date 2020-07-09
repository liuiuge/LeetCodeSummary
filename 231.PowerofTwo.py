# -*- coding:utf8 -*- 

# https://leetcode.com/problems/power-of-two/

class Solution:
    def isPowerOfTwo(self, n: int) -> bool:

        return n > 0 and n&(n-1) == 0

if __name__ == "__main__":
    sol = Solution()
    print(sol.isPowerOfTwo(1))
    print(sol.isPowerOfTwo(16))
    print(sol.isPowerOfTwo(218))