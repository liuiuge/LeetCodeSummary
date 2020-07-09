# -*- coding:utf8 -*-

class Solution:
    def groupAnagrams(self, strs):
        resultlist, resultdict = [], {}
        for strelem in strs:
            cnt = 1
            key = str(sorted(list(strelem)))
            if resultdict.get(key) is None:
                resultdict[key] = [strelem]
            else:
                resultdict[key].append(strelem)
        for _, elemlist in resultdict.items():
            resultlist.append(elemlist)
        return resultlist

if __name__ == "__main__":
    a = ["eat", "tea", "tan", "ate", "nat", "bat"]
    sol = Solution()
    result = sol.groupAnagrams(a)
    print(result)