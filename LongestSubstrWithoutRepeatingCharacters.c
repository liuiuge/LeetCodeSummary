int lengthOfLongestSubstring(char* s) {
    int hashtab[128];
    for(int i=0; i<128; ++i)
        hashtab[i] = -1;
    int start = -1;
    int len = strlen(s),maxlen = 0, tmplen = 0;
    for(int i=0; i<len ; ++i)
    {
        if(hashtab[s[i]]>start)
            start = hashtab[s[i]];
        hashtab[s[i]] = i;
        tmplen = i-start;
        if(maxlen<tmplen)
            maxlen = tmplen;
    }
    return maxlen;
}
