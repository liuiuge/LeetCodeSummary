#include <stdio.h>

int distributeCandies(int* candies, int candiesSize) {
    int len = candiesSize/2;
    char arr[200002] = {0};
    int i = 0,cnt = 0;
    for(;i<candiesSize && cnt <candiesSize/2; ++i )
    {
        if(arr[candies[i]+100001])
            continue;
        else{
            arr[candies[i]+100001] = 1;
            ++cnt;
        }
        if(cnt >= len)
            break;
    }
    return cnt;
}
