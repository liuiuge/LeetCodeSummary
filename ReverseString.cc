

#include <iostream>
#include <algorithm>
using std::reverse;
class Solution {
public:
    string reverseString(string s) {
        reverse(&s[0],&s[s.length()]);
        return s;
    }
};
