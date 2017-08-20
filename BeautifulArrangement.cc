

#include <iostream>
#include <vector>
#include <algorithm>
using std::vector;
using std::swap;
class Solution {
public:
    int countArrangement(int N) {
        if(N <= 0) return 1;
        vector<int> vs;
        for(int i=0;i<N;i++)
            vs.push_back(i+1);
        return counter(N,vs);
    }
    int counter(int n, vector<int> & vs)
    {
        if(n <= 0) return 1;
        int ans = 0;
        for(int i = 0; i < n; i++)
        {
            if( 0 == vs[i] % n || 0 == n % vs[i] )
            {
                swap(vs[i],vs[n-1]);
                ans += counter(n-1, vs);
                swap(vs[i],vs[n-1]);
            }
        }
        return ans;
    }
};
