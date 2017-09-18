class Solution {
public:
    int maxlen = 1, start = 0;
    void expand(string s, int low, int high){
        while(high<s.size() && low >= 0 && s[high] == s[low])
        {
            high++;
            low--;
        }
        if(high-low - 1>maxlen){
            start = low + 1;
            maxlen = high-low-1;
        }
    }
    string longestPalindrome(string s) {
        int len = s.size();
        if(len<2) return s;
        for(int i=0; i<len-1; ++i)
        {
            expand(s,i,i);
            expand(s,i,i+1);
        }
        return s.substr(start, maxlen);
    }
};
