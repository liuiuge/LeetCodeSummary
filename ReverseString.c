#include <stdio.h>
#include <assert.h>
#include <string.h>

char* reverseString(char* s) {
    assert(s != NULL);
    char * beg = s, *end = s+strlen(s)-1, *tmp = s;
    while(beg <end)
    {
        char temp = *beg;
        *beg = *end;
        *end = temp;
        ++beg;
        --end;
    }
    return tmp;
}
