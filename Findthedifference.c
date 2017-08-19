#include <stdio.h>
#include <string.h>

char findTheDifference(char* s, char* t) {
    char ans=0;
    for(int i = 0 ; s[i]!='\0'; i ++)ans^=(s[i]^t[i]);
    ans^=t[strlen(t)-1];
    return ans;
}
