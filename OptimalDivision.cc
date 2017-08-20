

#include <iostream>
#include <string>
#include <vector>
using std::string;
using std::to_string;
using std::vector;

class Solution {
public:
    string optimalDivision(vector<int>& nums) {
        if(!nums.size()) return string();
        string ans = to_string(nums[0]);
        if(1 == nums.size()) return ans;
        if(2 == nums.size()) return ans+"/"+to_string(nums[1]);
        ans += "/(" + to_string(nums[1]);
        for(size_t i =2; i< nums.size(); ++i)
            ans += "/" +to_string(nums[i]);
        ans += ")";
        return ans;
    }
};
