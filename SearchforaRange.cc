class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> ret(2,-1);
        if(nums.size()<1) return ret;
        auto low = lower_bound(nums.begin(),nums.end(),target);
        if(low == nums.end() || *low != target) return ret;
        else{
            auto up = upper_bound(nums.begin(),nums.end(),target);
            ret[0] = low - nums.begin();
            if(up != low)
                ret[1] = up - nums.begin() - 1;
            else
                ret[1] == up - nums.begin();
            return ret;
        }
    }
};
