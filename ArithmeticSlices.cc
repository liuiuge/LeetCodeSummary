

//Author:liuiuge(1220595883@qq.com

#include <iostream>
class Solution {
public:
   int numberOfArithmeticSlices(vector<int>& A) {
        int curr=0,sum=0;
        for(int i=2;i<A.size();i++)
        {
            if(A[i]-A[i-1]==A[i-1]-A[i-2])
            {
                curr++;
                sum+=curr;
            }
            else
            {
                curr=0;
            }
        }
        return sum;
    }
};
