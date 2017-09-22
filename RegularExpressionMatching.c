bool isMatch(char* s, char* p) {
    int m=strlen(s), n = strlen(p), ret = 0;
    int **dp = (int**)calloc(sizeof(int*),m+1);
    for(int i=0;i<=m; i++){
        dp[i] = (int*)calloc(sizeof(int),n+1);
    }
    dp[0][0] = true;
    for(int i = 0; i < m+1; i++) {
        for(int j = 1; j < n+1; j++) {
            if(p[j-1] != '*') {
                dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.');
            }
            else {
                dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.')); 
            }
        }
    }
    ret = dp[m][n];
    for(int i=0;i<=m;++i)
        free(dp[i]);
    free(dp);
    dp = NULL;
    return ret;    
}
