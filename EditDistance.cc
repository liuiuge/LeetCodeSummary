class Solution {
public:
    int minDistance(string word1, string word2) {
        int lenA = word1.size(), lenB=word2.size();
        int **ld = new int *[lenA+1];
        for(int i=0; i<=lenA; ++i)
            ld[i] = new int[lenB+1];
        for(int i=0; i<lenA; ++i)
            ld[i][lenB] = lenA-i;
        for(int j=0; j<lenB; ++j)
            ld[lenA][j] = lenB-j;
        ld[lenA][lenB] = 0;
        for(int i=lenA-1; i>=0; --i)
            for(int j=lenB-1; j>=0; --j)
            {
                if(word1[i] == word2[j])
                    ld[i][j] = ld[i+1][j+1];
                else
                    ld[i][j] = min(min(ld[i][j+1],ld[i+1][j]),ld[i+1][j+1])+1;
            }
        int dist= ld[0][0];
        for(int i=0; i<lenA+1;++i)
            delete []ld[i];
        delete []ld;
        return dist;
    }
};
