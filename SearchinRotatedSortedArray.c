int search(int* nums, int numsSize, int target) {
    int low = 0, high = numsSize - 1,mid;
    while(low<=high){
        mid = (high+low)/2;
        if(nums[mid] == target) return mid;
        else 
            if(nums[low] <= nums[mid]){
                if(target>nums[mid]){
                    low = mid + 1;
                }else{
                    if(target >= nums[low]){
                        high = mid -1;
                    }else{
                        low = mid + 1;
                    }
                }
            }else{
                if(target < nums[mid]){
                    high = mid -1;
                }else{
                    if(target <= nums[high]){
                        low = mid + 1;
                    }else{
                        high = mid -1;
                    }
                }
            }
    }
    return -1;
}
