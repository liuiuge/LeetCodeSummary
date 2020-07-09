# -*- coding: utf8 -*-

class Solution:
    def validPalindrome1(self, s: str) -> bool:
        l, r = 0, len(s) - 1
        while l < r:
            # print(s[l], s[r])
            if s[l] != s[r]:
                a, b, eqa = l, r, True
                while a+1 <= b:
                    # print(s[a+1], s[b])
                    if s[a+1] != s[b]:
                        eqa =False
                        break
                    a += 1
                    b += 1
                if eqa:
                    return eqa
                c, d = l, r
                while c + 1 <= d:
                    # print(s[c], s[d-1])
                    if s[c] != s[d-1]:
                        return False
                    c += 1
                    d -= 1
                return True
            l += 1
            r -= 1
        return True

    def validPalindrome(self, s: str) -> bool:
        if s[::-1] == s:
            return True
        l, r = 0, len(s) - 1
        while l < r:
            # print(s[l], s[r])
            if s[l] != s[r]:
                one, two = s[l+1:r+1], s[l:r]
                return one == one[::-1] or two == two[::-1]
            l += 1
            r -= 1
        return True
                
            
        
if __name__ == "__main__":
    sol = Solution()
    print(sol.validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))