#include <stdio.h>
#include <stdlib.h>
#include <math.h>

//Author:liuiuge(1220595883@qq.com)
int* findDisappearedNumbers(int* nums, int numsSize, int* returnSize) {
    for(int i=0; i<numsSize; ++i)
    {
        int index = abs(nums[i]-1);
        if(nums[index]>0)
            nums[index] = -nums[index];
    }
    int cnt =0;
    for(int i=0; i<numsSize; ++i)
    {
        if(nums[i]<0)
            ++cnt;
    }
    *returnSize = cnt;
    int *ret = (int *)malloc(sizeof(int)*cnt), j=0;
    for(int i=0; i<numsSize ; ++i)
    {
        if(nums[i]>0)
            ret[j++] = i+1;
    }
    return ret;
}
