
#include <stdlib.h>
#include <iostream>
#include <string>
using std::string;
using std::pair;
using std::to_string;
class Solution {
public:
    string complexNumberMultiply(string a, string b) {
        pair<int,int> ai,bi;
        int n1,n2;
        string result;

        ai=getints(a);
        bi=getints(b);

        n1=ai.first*bi.first-ai.second*bi.second;
        n2=ai.first*bi.second+ai.second*bi.first;

        result += to_string(n1);
        result += "+";
        result += to_string(n2);
        result += "i";
        return result;

    }

private:
    pair<int,int> getints(string a){
        size_t n;
        pair<int,int>ai;

        for(n=0;n<a.size();n++){
            if(a[n]=='+')break;
        }
        ai.first=atoi(a.substr(0,n).c_str());
        ai.second=atoi(a.substr(n+1,a.size()).c_str());

        return ai;
    }
};
