#include <stdio.h>
#include <string.h>
int findLUSlength(char* a, char* b) {
    int n = strlen(a), m = strlen(b);
    if(!strcmp(a,b))
        return -1;
    if(m == n)
        return n;
    return m>n?m:n;
}
