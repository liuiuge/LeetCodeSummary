#include <stdio.h>
#include <stdlib.h>

int* countBits(int num, int* returnSize) {
    int *ret = (int *)calloc(sizeof(int),(num+1));
    *returnSize = num+1;
    for(int i=1;i<=num;++i)
        ret[i] = ret[i>>1]+ i%2;
    return ret;
}
