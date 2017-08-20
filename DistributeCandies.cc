

#include <iostream>
#include <vector>
#include <unordered_set>
#include <algorithm>
using std::min;
using std::vector;
using std::unordered_set;
class Solution {
public:
    int distributeCandies(vector<int>& candies) {
        unordered_set<int> kinds;
        for (int kind : candies) {
            kinds.insert(kind);
        }
        return min(kinds.size(), candies.size() / 2);
    }
};
