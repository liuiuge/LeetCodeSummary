# -*- coding:utf8 -*-

class Solution:
    def twoSum(self, numbers, target):
        l, r = 0, len(numbers) - 1
        while l < r:
            while numbers[l] + numbers[r] < target:
                l += 1
            if numbers[l] + numbers[r] == target:
                return [l+1, r+1]
            r -= 1
        return []


if __name__ == "__main__":
    numbers = [2, 7, 11, 15]
    target = 9
    sol = Solution()
    print(sol.twoSum(numbers, target))