int removeElement(int* nums, int numsSize, int val) {
    int i = 0; 
    int n = numsSize - 1;   
    while (i <= n) {  
        if (nums[i] == val){           
            nums[i] = nums[n--]; 
        }  else { 
            i++; 
        }
     } 
     return n + 1;     
}
