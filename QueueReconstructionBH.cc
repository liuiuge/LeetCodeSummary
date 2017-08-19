

#include <iostream>
#include <vector>
#include <algorithm>
using std::vector;
using std::pair;
using std::sort;
class solution
{
public:
    vector<pair<int, int>> reconstructQueue(vector<pair<int, int>>& people) {
        sort(people.begin(), people.end(), [](const pair<int,int> &a, const pair<int,int> & b)
             {
             return (a.first>b.first || (a.first==b.first && a.second<b.second));
             });
        vector<pair<int,int>> ret;
        for(auto a:people)
            ret.insert(ret.begin()+a.second,a);
        return ret;
    }
};
