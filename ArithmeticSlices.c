#include <stdio.h>

int numberOfArithmeticSlices(int* A, int ASize) {
    int cnt=0,sum =0;
    for(int i=2; i<ASize;++i)
    {
        if(A[i]-A[i-1] == A[i-1]-A[i-2])
        {
            sum+=++cnt;
        }else
            cnt=0;
    }
    return sum;
}
