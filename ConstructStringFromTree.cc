

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <string>
using std::to_string;
using std::string;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
    public:
    void fun(TreeNode* t,string &s)
    {
        if(t)
        {
            int x=t->val;
            s+=to_string(x);
            if(!t->left && !t->right)
                return;
            s+='(';
            fun(t->left,s);
            s+=')';
            if(t->right)
            {
                s+='(';
                fun(t->right,s);
                s+=')';
            }

        }
    }
    string tree2str(TreeNode* t) {
        string s="";
        fun(t,s);
        return s;
    }
};
