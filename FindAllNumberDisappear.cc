
#include <math.h>
#include <iostream>
#include <vector>

using std::vector;
class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        for(size_t i=0; i<nums.size(); ++i)
        {
            nums[abs(nums[i]-1)] = -abs(nums[abs(nums[i]-1)]);
        }
        vector<int> ret;
        for(size_t i=0; i<nums.size(); ++i)
        {
            if(nums[i]>0)
                ret.push_back(i+1);
        }
        return ret;
    }
};
