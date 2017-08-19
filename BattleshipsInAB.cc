

#include <iostream>
#include <vector>
using std::vector;
class Solution {
public:
    int countBattleships(vector<vector<char>>& board) {
        if(board.empty())
            return 0;
        int cnt=0;
        int r = board.size();
        int c = board[0].size();
        for(int i =0; i < r; ++i)
            for(int j = 0; j < c; ++j)
                if(board[i][j] == 'X' && (i==0||board[i-1][j]=='.') && (board[i][j-1] == '.' || j==0))
                    cnt++;
        return cnt;
    }
};
