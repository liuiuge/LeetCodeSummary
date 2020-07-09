# -*- coding: utf8 -*-

class Solution:
    def isPalindrome(self, s: str) -> bool:
        if not s:
            return True
        l, r = 0, len(s) - 1
        while l < r:
            while not s[l].isalnum() and l < r:
                l += 1
            while not s[r].isalnum() and l < r:
                r -= 1
            if s[l].upper() != s[r].upper():
                return False
            l += 1
            r -= 1
        return True

if __name__ == "__main__":
    sol = Solution()
    print(sol.isPalindrome("A man, a plan, a canal: Panama"))
