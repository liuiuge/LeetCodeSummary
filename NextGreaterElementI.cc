

#include <iostream>
#include <vector>
#include <stack>
#include <unordered_map>
using std::unordered_map;
using std::vector;
using std::stack;
class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& findNums, vector<int>& nums) {
        vector<int> temp;
        stack<int> s1;
        unordered_map<int,int> ma1;
        temp.resize(findNums.size());
        for(auto elem : nums)
        {
            while(!s1.empty()&& elem > s1.top())
            {
                ma1[s1.top()] = elem;
                s1.pop();
            }
            s1.push(elem);
        }
        while(!s1.empty())
        {
            ma1[s1.top()] = -1;
            s1.pop();
        }
        for(size_t i =0; i < findNums.size(); ++i)
            temp[i] = ma1[findNums[i]];
        return temp;
    }
};
