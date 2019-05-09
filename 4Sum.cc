class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        size_t len = nums.size();
        if(len < 4)
        {
            return vector<vector<int> > ();
        }
        vector<vector<int> > ret;
        set<vector<int> > sret;
        sort(nums.begin(), nums.end());
        for(size_t i = 0; i < len - 3; ++i)
        {
            if (i > 0 && nums[i] == nums[i-1]) 
                continue;
            if (nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target)
                break;
            if(nums[i] + nums[len-3] + nums[len-2] + nums[len-1] < target)
                continue;
            
            for(size_t j = i + 1; j < len - 2; ++j)
            {
                if (j > i+1 && nums[j] == nums[j-1]) 
                    continue;
                if (nums[i] + nums[j+1] + nums[j+2] + nums[j] > target)
                    break;
                if(nums[i] + nums[j] + nums[len-2] + nums[len-1] < target)
                    continue;
                
                size_t begin = j + 1;
                size_t end = len - 1;
                while(begin < end)
                {
                    int sum = nums[i] + nums[j] + nums[begin] + nums[end];
                    if (sum == target) 
                    {
                        vector<int> tmpret = {nums[i],  nums[j], nums[begin], nums[end]};
                        sret.insert(tmpret);
                        ++ begin;
                        --end;
                    }
                    else if (sum < target)
                    {
                        ++ begin;
                    }
                    else if (sum > target)
                    {
                        -- end;
                    }
                }
            }
        }
        
        if (sret.size() < 1)
        {
            return vector<vector<int> > ();
        }
        ret.reserve(sret.size());
        for(auto it=sret.begin(); it != sret.end(); ++it)
        {
            ret.push_back(*it);
        }
        return ret;
    }
};
