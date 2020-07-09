# -*- coding: utf8 -*-

class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        dp = [[0] * (len(text1) + 1) for _ in range(len(text2) + 1)]
        for i, c in enumerate(text2):
            for j, d in enumerate(text1):
                dp[i + 1][j + 1] = 1 + dp[i][j] if c == d else max(dp[i][j + 1], dp[i + 1][j])
        [print(row) for row in dp]
        return dp[-1][-1]

if __name__ == "__main__":
    sol = Solution()
    # print(sol.longestCommonSubsequence("abcde", "ace"))
    # print(sol.longestCommonSubsequence("bl", "yby"))
    # print(sol.longestCommonSubsequence("hofubmnylkra","pqhgxgdofcvmr"))
    print(sol.longestCommonSubsequence("papmretkborsrurgtina","nsnupotstmnkfcfavaxgl"))
