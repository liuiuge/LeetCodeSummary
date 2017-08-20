

//Author:liuiuge(1220595883@qq.com

#include <stdlib.h>
#include <iostream>
#include <string>
using std::string;
class Solution {
public:
    bool detectCapitalUse(string word) {
        if (word.length() <= 1) return true;
        bool firstcap = isupper(word[0]);
        bool secondcap = isupper(word[1]);
        if (!firstcap && secondcap) return false;
        for (size_t c = 2; c < word.length(); c++) {
            if (!firstcap && isupper(word[c])) return false;
            if (secondcap != !!isupper(word[c])) return false;
        }
        return true;
    }
};
