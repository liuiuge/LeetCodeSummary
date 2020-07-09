# -*- coding:utf8 -*-

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) == len(t):
            if s == t:
                return True
            dicts, dictt = {}, {}
            for x in s:
                dicts[x] = dicts.get(x, 0) + 1
            for y in t:
                dictt[y] = dictt.get(y, 0) + 1
            for k, v in dicts.items():
                if dictt.get(k) != v:
                    return False
            return True

        return False

if __name__ == "__main__":
    sol = Solution()
    print(sol.isAnagram("rat", "car"))
    print(sol.isAnagram("rat", "tar"))