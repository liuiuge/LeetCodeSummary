class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if(strs.size() == 0) return "";
        if(strs.size() == 1) return strs[0];
        int i, j;
        int index = 0;
        bool isE = false;
        bool isOver = false;
        for(i = 0; i < strs[0].length(); i++){
            isE = true;
            for(j = 1; j < strs.size(); j++){
                if(i >= strs[j].length()){
                    isOver = true;
                    isE = false;
                    break;
                }
                if(strs[0][i] == strs[j][i]){
                    continue;
                } else {
                    isE = false;
                    break;
                }
            }
            if(isOver) break;
            if(isE) {
                index++;
            } else {
                break;
            }
        }
        string s(index, 'a');
        for(i = 0; i < index; i++){
            s[i] = strs[0][i];
        }
        return s;
    }
};
