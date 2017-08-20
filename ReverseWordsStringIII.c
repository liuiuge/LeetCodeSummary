#include <stdio.h>

//Author:liuiuge(1220595883@qq.com)
char* reverseWords(char* s) {
    char *beg = s;
    char *end = s;
    char temp = 0;
    while(*beg != '\0')
    {
        if(*beg == ' ')
        {
            ++beg;
            ++end;
        }
        else if(*end == '\0' || *end == ' ')
        {
            char *ps = beg, *pe = end-1;
            while(ps <pe)
            {
                temp = *ps;
                *ps = *pe;
                *pe = temp;
                ps++;
                pe--;
            }
            beg = end;
        }
        else
            end++;
    }
    return s;
}
