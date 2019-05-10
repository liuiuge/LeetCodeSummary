class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 0:
            return 0
        if n <= 1:
            return 1
        s1, s2 = 0, 1
        for i in xrange(1, n+1):
            t = s2
            s2 = s2 + s1
            s1 = t
        return s2
