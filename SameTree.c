/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

bool GetVal(struct TreeNode *p, struct TreeNode* q) {
    if (p == NULL) 
    {
        if(q == NULL)
        {
            return true;
        }
        else
        {
            return false;
        }
    }
    else {
        if (q == NULL)
        {
            return false;
        }
        else if (p -> val != q->val)
        {
            return false;
        }
        return GetVal(p->left, q->left) & GetVal(p->right, q->right);
    }
}

bool isSameTree(struct TreeNode* p, struct TreeNode* q){
    if (p == NULL && q == NULL) return true;
    return GetVal(p, q);
}
