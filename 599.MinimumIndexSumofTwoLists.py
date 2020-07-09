# -*- coding:utf8 -*-

class Solution:
    def findRestaurant(self, list1, list2):
        x = dict()
        for i in range(len(list1)):
            x[list1[i]] = i
        ret, cnt = [], 2000
        for j in range(len(list2)):
            if list2[j] in x:
                if x[list2[j]] + j < cnt:
                    cnt = x[list2[j]] + j
                    ret = [list2[j]]
                elif  x[list2[j]] + j == cnt:
                    ret.append(list2[j])
        return ret


if __name__ == "__main__":
    sol = Solution()
    a = ["Shogun", "Tapioca Express", "Burger King", "KFC"]
    b = ["KFC", "Shogun", "Burger King"]
    print(sol.findRestaurant(a,b))