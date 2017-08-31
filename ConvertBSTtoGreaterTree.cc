class Solution {
public:
    int cnt = 0;
    void valuecalc(TreeNode* root)
    {
        if(!root) return;
        if(root->right) valuecalc(root->right);
        cnt += root->val;
        root->val = cnt;
        if(root->left) valuecalc(root->left);
    }
    TreeNode* convertBST(TreeNode* root) {
        valuecalc(root);
        return root;
    }
};
