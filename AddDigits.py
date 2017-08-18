#!/usr/bin/env python
# coding=utf-8
class Solution(object):
    def addDigits(self, num):
        if num==0:
            return 0;
        else:
            return (num-1)%9+1
