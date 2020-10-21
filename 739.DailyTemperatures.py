# -*- coding: utf-8 -*-

'''
int[] result = new int[T.Length];
for(int i = 0;i<T.Length;i++)
{
    int key = T[i];
    int count = 1;
    for(int j = i + 1;j<T.Length;j++)
    {
        if (key < T[j])
        {
            result[i] = count;
            break;
        }
        else
        {
            count++;
        }
    }
}
return result;
}
'''

class Solution:
    def dailyTemperatures(self, T):
        ret = [0 for _ in T]
        i = 0
        while i < len(T) - 1:
            if i > 1:
                if T[i] == T[i-1] and ret[i-1] > 0:
                    ret[i] = ret[i-1] - 1
                    i += 1
                    continue
            j = 1
            while i+j < len(T):
                if T[i+j] > T[i]:
                    ret[i] = j
                    break
                j += 1
            i += 1
        return ret

if __name__ == "__main__":
    pass