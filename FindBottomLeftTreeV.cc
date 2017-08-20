

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <queue>
using std::queue;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
public:
    int findBottomLeftValue(TreeNode* root) {
        if(!root) return 0;
        int ret = root->val;
        queue<TreeNode*> s1;
        s1.push(root);
        while(!s1.empty())
        {
            int sz = s1.size();
            for(int i =0; i<sz; ++i)
            {
                TreeNode* tmp = s1.front();
                s1.pop();
                if(i==0) ret = tmp->val;
                if (tmp->left) s1.push(tmp->left);
                if (tmp->right) s1.push(tmp->right);
            }
        }
        return ret;
    }
};
