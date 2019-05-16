/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        queue<TreeNode *> q1;
        vector<vector<int>> v1;
        if (root == NULL){
            return v1;
        }
        TreeNode * cur = root;
        int level_length = 1;
        q1.push(cur);
        while(!q1.empty())
        {
            vector<int> tmp;
            int i_tmp =0;
            for(int i=0; i < level_length; ++i)
            {
                TreeNode * p = q1.front();
                tmp.push_back(p->val);
                q1.pop();
                if(p->left)
                {
                    q1.push(p->left);
                    ++i_tmp;
                }
                if(p->right)
                {
                    q1.push(p->right);
                    ++i_tmp;
                }
            }
            level_length = i_tmp;
            v1.push_back(tmp);
        }
        reverse(v1.begin(), v1.end());
        return v1;
    }
};
