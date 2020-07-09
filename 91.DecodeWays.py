# -*- coding:utf8 -*-

# https://leetcode.com/problems/decode-ways/

class Solution:
    def numDecodings(self, s: str) -> int:
        if not s:
            return 0
			
        if len(s)==1:
            return 1 if s[0]!="0" else 0
        
        np2 = 1 if s[-1]!="0" else 0
		
        # check last but one character
        if (s[-2]=="1") or ((s[-2]=="2") and (s[-1] not in ("7","8","9"))):
            np1 = 1+ np2
        else:
            np1 = np2 if s[-2]!="0" else 0

        for i in range(len(s)-3,-1,-1):
            n = 0
            if (s[i]=="1") or ((s[i]=="2") and (s[i+1] not in ("7","8","9"))):
                n = np2
            if s[i]!="0":
                n+= np1
            np1,np2 = n,np1
                               
        return np1 


if __name__ == "__main__":
    sol = Solution()
    print(sol.numDecodings("2226"))