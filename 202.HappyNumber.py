# -*- coding:utf8  -*-

class Solution:
    def isHappy(self, n: int) -> bool:
        result = []
        while n != 1:
            cnt = 0
            for i in str(n):
                cnt += int(i) * int(i)
            if cnt in result:
                return False
            result.append(cnt)
            n = cnt
            
        return True

if __name__ == "__main__":
    sol = Solution()
    print(sol.isHappy(19))
