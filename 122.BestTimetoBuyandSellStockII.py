# -*- coding:utf8 -*-

# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

# pos_open = False
#         for i in range(1,len(prices)):
#             if prices[i-1]<prices[i] and not pos_open:
#                 buy_i = i-1
#                 pos_open = True
#             elif prices[i-1]>prices[i] and pos_open:
#                 profit += prices[i-1] - prices[buy_i]
#                 pos_open = False
                
#             if pos_open and i == len(prices)-1:
#                 profit += prices[i] - prices[buy_i]
#                 pos_open = False

class Solution:
    def maxProfit(self, prices) -> int:
        if len(prices) < 2:
            return 0
        brise = False
        buy_i = 0
        cnt = 0
        for i in range(1, len(prices)):
            
            if prices[i-1]<prices[i] and not brise:
                buy_i = i-1
                brise = True
            elif prices[i-1]>prices[i] and brise:
                cnt += prices[i-1] - prices[buy_i]
                brise = False
                
            if brise and i == len(prices)-1:
                cnt += prices[i] - prices[buy_i]
                brise = False


        return cnt

if __name__ == "__main__":
    sol = Solution()
    print(sol.maxProfit([7,6,4,3,1]))
    print(sol.maxProfit([7,1,5,3,6,4]))
