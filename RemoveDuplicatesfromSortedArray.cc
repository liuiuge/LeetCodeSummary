class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        nums.resize(distance(nums.begin(),unique(nums.begin(),nums.end())));
        return nums.size();
    }
};
