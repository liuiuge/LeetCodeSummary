#include <stdio.h>

int addDigits(int num) {
    int ret = 0;
    while(num)
    {
        ret += num%10;
        num/=10;
    }
    return ret<10?ret:addDigits(ret);
}
