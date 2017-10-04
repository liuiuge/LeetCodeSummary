/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode* plist;
plist merge2list(plist l1, plist l2){
    if(!l1) return l2;
    if(!l2) return l1;
    if(l1->val <= l2->val) {
        l1->next = merge2list(l1->next,l2);
        return l1;
    }else{
        l2->next = merge2list(l1,l2->next);
        return l2;
    }
}
struct ListNode* mergeKLists(struct ListNode** lists, int listsSize) {
    if(lists == NULL) return NULL;
    while(listsSize > 1){
        for(int i=0;i<listsSize/2;++i)
            lists[i] = merge2list(lists[i],lists[listsSize-1-i]);
        listsSize = (listsSize+1)/2;
    }
    return lists[0];
}
