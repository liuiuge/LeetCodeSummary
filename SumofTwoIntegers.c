#include <stdio.h>

//Author:liuiuge(1220595883@qq.com)
int add(int a, int b)
{
    if(b==0) return a;
    int sum =a^b;
    int carry = (a&b)<<1;
    return add(sum,carry);
}


int getSum(int a, int b) {
    return add(a,b);
}
