#!/usr/bin/env python
# coding=utf-8
class Solution(object):
    def islandPerimeter(self, grid):
        m,n = len(grid),len(grid[0])
        cnt,neighbor=0,0
        for i in xrange(m):
            for j in xrange(n):
                if(grid[i][j]):
                    cnt+=1
                    if i<m-1 and grid[i+1][j]:
                         neighbor+=1
                    if j<n-1 and grid[i][j+1]:
                         neighbor+=1
        return 4*cnt-2*neighbor
