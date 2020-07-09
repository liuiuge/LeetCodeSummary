# -*- coding:utf8 -*-

# https://leetcode.com/problems/ugly-number-ii/

def isUglyNum(n) -> bool:
    if n %2 != 0 and n%5!=0 and n%3!=0:
        return False
    while n % 2 == 0:
        n /= 2
    while n % 3 == 0:
        n /= 3
    while n % 5 == 0:
        n /= 5
    if n == 1:
        return True
    return False

class Solution: 
    def nthUglyNumber1(self, n: int) -> int:
        if n < 2:
            return 1
        cnt, i, ret = 1, 1, [], 
        while cnt < n:
            if isUglyNum(i):
                cnt += 1
                ret.append(i)
            i += 1
        return ret[-1]

    def nthUglyNumber(self, n: int) -> int:
        i1, i2, i3 = 0, 0, 0
        if n < 2:
            return 1
        cnt, ret = 1, [0 for i in range(n)]
        ret[0] = 1
        while cnt < n:
            val = min(ret[i1]*2, ret[i2]*3, ret[i3] *5)
            ret[cnt] = val
            if val == ret[i1]*2:
                i1 += 1
            if val == ret[i2]*3:
                i2 += 1
            if val == ret[i3] *5:
                i3 += 1
            cnt += 1
            
        return ret[-1]

            
        

if __name__ == "__main__":
    sol = Solution()
    print(sol.nthUglyNumber(1000))