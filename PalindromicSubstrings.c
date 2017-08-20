#include <stdio.h>
#include <string.h>

//Author:liuiuge(1220595883@qq.com)
int ext(char *s, int left, int right)
{
    int ans = 0, len = strlen(s);
    while(left>=0 && right<len && s[left] == s[right])
    {
        ++right;
        --left;
        ++ans;
    }
    return ans;
}
int countSubstrings(char* s) {
    int ans=0, len= strlen(s);
    for(int i=0; i<len; ++i)
    {
        ans+= ext(s,i,i);
        ans+= ext(s,i,i+1);
    }
    return ans;
}
