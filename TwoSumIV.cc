

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <vector>
#include <queue>
using std::vector;
using std::queue;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        if(!root) return 0;
        vector<int> a;
        queue<TreeNode*> s1;
        s1.push(root);
        while(!s1.empty())
        {
            int sz = s1.size();
            for(int i =0;i<sz;i++)
            {
                TreeNode * t = s1.front();
                a.push_back(t->val);
                if(t->left) s1.push(t->left);
                if(t->right) s1.push(t->right);
                s1.pop();
            }
        }
        for(size_t i =0; i< a.size()-1;++i)
        {
            for(size_t j = i+1; j < a.size(); ++j )
                if(a[i]+a[j] == k)
                    return 1;
        }
        return 0;
    }
};
