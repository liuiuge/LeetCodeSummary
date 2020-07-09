# -*- coding:utf8 -*-
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

class Solution:
    def maxProfit(self, prices) -> int:
        if not len(prices):
            return 0
        minprofit, maxprofit =prices[0], 0
        for i in range(len(prices)):
            if prices[i] < minprofit:
                minprofit = prices[i]
            if prices[i] - minprofit > maxprofit:
                maxprofit = prices[i] - minprofit

        return maxprofit
                    

if __name__ == "__main__":
    sol = Solution()
    print(sol.maxProfit([7,6,4,3,1]))
    print(sol.maxProfit([7,1,5,3,6,4]))