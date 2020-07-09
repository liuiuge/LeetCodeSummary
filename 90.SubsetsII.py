# -*- coding:utf8 -*-

# https://leetcode.com/problems/subsets-ii/

'''
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.
Example:
Input: [1,2,2]
Output:
[ [2], [1], [1,2,2], [2,2], [1,2], []]

sample:
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        self.res = []
        self.helper([], 0, nums)
        return self.res
        
    def helper(self, path, depth, nums):
        if depth >= len(nums):
            if path not in self.res:
                self.res.append(path)
            return
        self.helper(path, depth+1, nums)
        self.helper(path+[nums[depth]], depth+1, nums)

'''

class Solution:
    def subsetsWithDup(self, nums):
        ret = [[]]
        for i in range (1 << len(nums)):
            j, k, tmp, same = i, 0, [], 0
            while j:
                if j & 1:
                    tmp.append(nums[k])
                j >>= 1
                k += 1
            for x in ret:
                if sameList(x, tmp):
                    same = 1
                    break
            if not same:
                ret.append(tmp)
        return ret

def sameList(nums1, nums2):
    if len(nums1) == len(nums2):
        snum1 = sorted(nums1)
        snum2 = sorted(nums2)
        for i in range(len(nums1)):
            if snum1[i] != snum2[i]:
                return False
        return True
    return False

if __name__ == "__main__":
    sol = Solution()
    ret = sol.subsetsWithDup([1,2,3])
    print(ret)