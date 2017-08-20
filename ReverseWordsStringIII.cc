

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <algorithm>
using std::reverse;
class Solution {
public:
    string reverseWords(string s) {
        size_t front = 0;
        for(int i = 0; i <= s.length(); ++i)
            if(i == s.length() || s[i] == 0x20){
                reverse(&s[front], &s[i]);
                front = i + 1;
            }
        return s;
    }
};
