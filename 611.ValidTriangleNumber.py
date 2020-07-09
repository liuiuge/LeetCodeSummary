# -*- coding: utf8 -*-

def isfit(a):
    return a[0]+a[1] > a[2] and a[1] + a[2] > a[0] and a[0] + a[2] > a[1]

class Solution:
    def triangleNumber1(self, nums) -> int:
        length = len(nums)
        cnt = 0
        for i in range(length-2):
            for j in range(i+1, length-1):
                for k in range(j+1, length):
                    a = [nums[i], nums[j], nums[k]]
                    if isfit(a):
                        cnt += 1
                        print(a)
        return cnt

    def triangleNumber(self, nums) -> int:
        c = 0
        n = len(nums)
        nums.sort()
        for i in range(n-1,1,-1):
            lo = 0
            hi = i - 1
            while lo < hi:
                if nums[hi]+nums[lo] > nums[i]:
                    # 最大的 加 最小的大于更大的
                    # 则最大的 加 最小到最大之间的数加最大都可以满足
                    c += hi-lo
                    hi -= 1
                else:
                    lo += 1
        return c

if __name__ == "__main__":
    sol = Solution()
    print(sol.triangleNumber([2,2,2,3]))