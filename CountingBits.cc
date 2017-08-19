

#include <iostream>
#include <vecotr>
using std::vector;

class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> ret(num+1, 0);  
        for(int i=1; i<=num; ++i)  
            ret[i] = ret[i>>1] + i%2;  
        return ret;
    }
};
