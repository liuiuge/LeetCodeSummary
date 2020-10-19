# -*- coding:utf-8 -*-

# https://leetcode.com/problems/sort-array-by-parity-ii/

class Solution(object):
    def sortArrayByParityII(self, A):
        i = 0
        while i < len(A):
            if (A[i]&1)^(i&1) == 0:
                i += 1
                continue
            else:
                j = i + 1
                while j < len(A) :
                    if (A[j]&1) ^ (j&1) > 0 and (A[j]&1) ^ (A[i]&1) > 0:
                        break
                    j += 2
                A[i], A[j] = A[j], A[i]
                i +=1
        return A

    def sortArrayByParityIIFast(self, A):
        j = 1
        for i in range(0, len(A), 2):
            if A[i] % 2:
                while A[j] % 2:
                    j += 2
                A[i], A[j] = A[j], A[i]
        return A

if __name__ == "__main__":
    sol = Solution()
    print(sol.sortArrayByParityII([648,831,560,986,192,424,997,829,897,843]))
