

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <vector>
using std::vector;

class Solution {
public:
    vector<vector<int>> matrixReshape(vector<vector<int>>& nums, int r, int c) {
        int h = nums.size();
        int w = nums[0].size();
        if(h*w != r*c) return nums;
        vector<vector<int>> new1(r,vector<int>(c,0));
        int m=0 , n=0;
        for(int i = 0; i < h; ++i)
        {
            for(int j = 0; j < w;++j)
            {

                new1[m][n++] = nums[i][j];
                if(n == c)
                {
                    ++m;
                    n = 0;
                }
            }
        }
        return new1;
    }
};
