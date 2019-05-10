class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        retList = []
        if len(a) < len(b):
            a, b = b, a
        rise = 0
        for i in range(1, len(a) + 1):
            aVal = int(a[-1 * i])
            bVal = 0
            if (len(b) - i) >= 0:
                bVal = int(b[-1 * i])
            tmp = aVal + bVal + rise
            if tmp == 0:
                retList.insert(0, '0')
                rise = 0
            elif tmp == 1:
                retList.insert(0, '1')
                rise = 0
            elif tmp == 2:
                retList.insert(0, '0')
                rise = 1
            elif tmp == 3:
                retList.insert(0, '1')
                rise = 1
        if rise == 1:
            retList.insert(0, '1')
        return ''.join(retList)
