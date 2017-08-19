#include <stdio.h>

//Author:liuiuge(1220595883@qq.com)
int islandPerimeter(int** grid, int gridRowSize, int gridColSize) {
    int cnt = 0, neighbor = 0;
    for(int i=0;i<gridRowSize; ++i)
        for(int j=0; j<gridColSize; ++j)
        {
            if(grid[i][j])
            {
                cnt++;
                if(i>0&&grid[i-1][j])
                    neighbor++;
                if(j>0&&grid[i][j-1])
                    neighbor++;
            }

        }
    return 4*cnt-2*neighbor;
}
