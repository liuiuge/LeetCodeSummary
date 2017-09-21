int compare(const void*p1, const void*p2){
    return *(int*)p1-*(int*)p2;
}
int threeSumClosest(int* nums, int numsSize, int target) {
    qsort(nums,numsSize,4,compare);
    int j;
    int k;
    int distance,abs;
    int lo_distance = INT_MAX;
    int sum;


    for(int i = 0; i < numsSize-2;i++){
        j = i+1;
        k = numsSize-1;
        while(j<k){
            distance = nums[i]+nums[j]+nums[k] - target;
            abs = (distance >=0)? distance:-distance;
            if (abs < lo_distance){
                sum = nums[i]+nums[j]+nums[k];
                lo_distance = abs;
            }

            if(distance > 0) k--;
            else if(distance < 0) j++;
            else return target;
        }
    }
    return sum;
}
