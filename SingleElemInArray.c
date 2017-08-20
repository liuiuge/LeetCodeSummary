#include <stdio.h>

int singleNonDuplicate(int* nums, int numsSize) {
    int num = 0;
    for(int i =0; i < numsSize; ++i)
        num ^= nums[i];
    return num;
}
