int searchInsert(int* nums, int numsSize, int target) {
    int mid = numsSize/2;
    while(1){
        if(target == nums[mid]) return mid;
        else if(target > nums[mid]){
            if(mid+1 == numsSize || nums[mid+1] >= target) 
                return mid+1;
            else 
                ++mid;
        }else{
            if(mid-1 == -1) 
                return 0;
            else if(nums[mid-1] < target) 
                return mid;
            else 
                --mid;
        }
    }
    return 0;
}
