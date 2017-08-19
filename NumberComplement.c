#include <stdio.h>

int findComplement(int num) {
    int i = 0,ret = 0,tmp = num;
    while((tmp/=2)!=0)
       ret += 1<<i++;
    ret+=1<<i;
    return ret-num;
}
