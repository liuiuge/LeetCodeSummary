#include <stdio.h>

int hammingDistance(int x, int y) {
    int z = x ^ y,cnt=0;
    for(int i=0;i<32;i++) cnt+=(z>>i)&1;
    return cnt;
}
