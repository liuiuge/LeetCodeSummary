bool judgeCircle(char* moves) {
    int len = strlen(moves);
    int m=0, n=0;
    for(int i=0; i<len; ++i)
    {
        switch(moves[i])
        {
            case'L':m++;break;
            case'R':m--;break;
            case'U':n++;break;
            case'D':n--;break;    
        }
    }
    return !m&&!n;
}
