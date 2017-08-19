#include <stdio.h>

int totalHammingDistance(int* nums, int numsSize) {
    int cnt = 0, ret = 0;
    for(int i = 0; i < 32; ++i)
    {
        cnt = 0;
        for(int j=0; j < numsSize ; ++j)
        {
            if((nums[j]>>i) & 1)
                cnt++;
        }
        ret += cnt*(numsSize-cnt);
    }
    return ret;
}
