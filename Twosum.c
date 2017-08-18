#include <stdio.h>
#include <stdlib.h>

int* twoSum(int* nums, int numsSize, int target) {
    int *ts = (int *)calloc(2,sizeof(int));
    int flag = 0;
    for(int i = 0; i<numsSize-1; ++i)
    {
        int tsa = target - nums[i];
        for(int j = i + 1; j < numsSize; ++j)
        {
            if(nums[j] == tsa){
                flag = 1;
                ts[0] = i;
                ts[1] = j;
                break;
            }
        }
        if(flag == 1)
            break;
    }
    if(flag == 1)
        return ts;
    return NULL;
}
