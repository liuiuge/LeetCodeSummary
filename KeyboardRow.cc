

//Author:liuiuge(1220595883@qq.com

#include <iostream>
class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        vector<int> v1(26);
        vector<string> dict = {"QWERTYUIOP", "ASDFGHJKL", "ZXCVBNM"},ret;
        int bit =0;
        for(auto & elem:dict)
        {
            for(char & c: elem)
                v1[c-'A'] = 1<<bit;
            ++bit;
        }
        for(auto & word:words)
        {
            int flag = 7;
            for(char & c:word)
            {
                if((flag &= v1[toupper(c)-'A']) == 0)
                    break;
            }
            if(flag)
                ret.push_back(word);
        }
        return ret;
    }
};
