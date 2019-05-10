class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        t, r = 0, x
        while t <= r:
            mid = (r-t) // 2 + t
            if mid * mid <= x and (mid+1)*(mid+1) > x:
                return mid
            elif mid * mid > x:
                r = mid
            else:
                t = mid + 1
