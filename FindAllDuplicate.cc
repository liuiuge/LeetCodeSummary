

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <vector>
using std::vector;
class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> new1;
        int len = nums.size();
        for(int i =0; i < len; ++i)
        {
            int index = abs(nums[i])-1;
            if(nums[index]<0)
                new1.push_back(index+1);
            nums[index] = -nums[index];
        }
        return new1;
    }
};
