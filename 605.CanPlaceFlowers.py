# -*- coding:utf-8 -*-

# https://leetcode.com/problems/can-place-flowers/

class Solution:
    def canPlaceFlowers(self, flowerbed, n) -> bool:
        if n < 1:
            return True
        if len(flowerbed) < 1:
            return False
        space, pre = 0, 0
        for i in range(len(flowerbed)):
            
            if flowerbed[i] > 0:
                continue
            else:
                print(i, flowerbed[i], pre, space)
                if i != len(flowerbed)-1 and flowerbed[i+1] > 0:
                    pre = 0
                    continue
                if i > 0 and flowerbed[i-1] > 0:
                    pre = 0
                    continue
                if pre > 0:
                    pre = 0
                    continue
                space += 1
                pre = 1
        print(space)
        if space >= n:
            return True
        return False

if __name__ == "__main__":
    sol = Solution()
    print(sol.canPlaceFlowers([1,0,0,0,1,0,0], 2))