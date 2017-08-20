

//Author:liuiuge(1220595883@qq.com

#include <iostream>
#include <queue>
#include <vector>
using std::queue;
using std::vector;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> avg;
        queue<TreeNode*> ts{{root}};
        while(!ts.empty())
        {
            int n = ts.size();
            double sum = 0;
            for(int i=0;i<n;++i)
            {
                TreeNode * tmp = ts.front();
                ts.pop();
                sum+= tmp->val;
                if(tmp->left)
                    ts.push(tmp->left);
                if(tmp->right)
                    ts.push(tmp->right);
            }
            avg.push_back(sum/n);
        }
        return avg;
    }
};

