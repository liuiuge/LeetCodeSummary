#include <stdio.h>

//Author:liuiuge(1220595883@qq.com)
int findMaxConsecutiveOnes(int* nums, int numsSize) {
    int Max=0,tempMax=0;
    int i=0;
    for(;i<numsSize;i++){
        if(*(nums+i)==1){
            if(++tempMax>Max)
                Max = tempMax;
        }
        else{
            tempMax=0;
        }      
    }
    return Max;
}    
