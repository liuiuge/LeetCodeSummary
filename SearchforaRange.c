int* searchRange(int* nums, int numsSize, int target, int* returnSize) {
    int *ret = (int*)malloc(8);
    ret[0] = -1;ret[1] = -1;
    *returnSize = 2;
    if(numsSize<1 || target<nums[0] || target>nums[numsSize-1]) return ret;
    int low =0, high = numsSize - 1, mid;
    while(low<=high){
        mid = high-(high-low)/2;
        if(nums[mid]>=target){
            high = mid - 1;
        }else{
            low = mid + 1;
        }
    }
    if(low > numsSize || nums[low] != target) return ret;
    ret[0] = low;
    low = 0; high = numsSize -1;
    while(low<=high){
        mid = high - (high-low)/2;
        if(nums[mid]<=target){
            low = mid + 1;
        }else{
            high = mid -1;
        }
    }
    ret[1] = high;
    return ret;
}
