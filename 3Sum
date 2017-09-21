class Solution {
public:
    vector<vector<int> > threeSum(vector<int> &nums) {
        vector<vector<int > > ret;
        sort(nums.begin(),nums.end());
        int len = nums.size();
        for(int i=0;i<len-2;++i){
            int front = i + 1, rear = len-1, target = -nums[i];
            while(front < rear){
                int sum = nums[front]+nums[rear];
                if(sum < target) front++;
                else if(sum > target) rear--;
                else{
                    vector<int> tmp(3,0);
                    tmp[0] = nums[i]; tmp[1] = nums[front]; tmp[2] = nums[rear];
                    ret.push_back(tmp);
                    while(front< rear && nums[front] == tmp[1]) front++;
                    while(front< rear && nums[rear] == tmp[2]) rear--;
                }
            }
            while(i+1<len-2 && nums[i+1] == nums[i]) ++i;
        }
        return ret;
    }
};
