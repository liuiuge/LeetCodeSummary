/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* deleteDuplicates(struct ListNode* head){
    if (head == NULL || head->next == NULL) {
        return head;
    }
    struct ListNode* cur = head;
    struct ListNode* pre = NULL;
    while(cur != NULL) {
        pre = cur;
        cur = cur ->next;
        while(cur!= NULL && cur->val == pre->val){
            pre ->next = cur -> next;
            cur = cur ->next;
        }
    }
    return head;
}
