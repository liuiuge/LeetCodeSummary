#include <stdio.h>
#include <stdlib.h>

//Author:liuiuge(1220595883@qq.com)
int comp(const void *a, const void *b)
{return *(int*)a-*(int*)b;}

int arrayPairSum(int* nums, int numsSize) {
    qsort(nums,numsSize,sizeof(int),comp);
    int sum =0;
    for(int i=0;i<numsSize/2;++i)
        sum += nums[2*i];
    return sum;
}
