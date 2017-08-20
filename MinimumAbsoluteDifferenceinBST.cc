

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <stack>
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    int getMinimumDifference(TreeNode* root) {
        int minDiff = INT_MAX, pre = -1;
        stack<TreeNode*> s1;
        TreeNode * p = root;
        while(p||!s1.empty())
        {
            while(p)
            {
                s1.push(p);
                p=p->left;
            }
            p = s1.top();
            s1.pop();
            if(pre!=-1) minDiff = min(p->val-pre,minDiff);
            pre = p->val;
            p= p->right;
        }
        return minDiff;
    }
};
