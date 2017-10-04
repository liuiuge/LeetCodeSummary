/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode* pList;
typedef struct ListNode Node;
pList reverseKGroup(pList head, int k) {
    if(head==NULL||k==1) return head;
    int num=0;
    pList preh = (pList)calloc(1,sizeof(Node));
    preh->next = head;
    pList cur = preh, nex, pre = preh;
    while(cur = cur->next) ++num;
    while(num>=k) {
        cur = pre->next;
        nex = cur->next;
        for(int i=1;i<k;++i) {
            cur->next=nex->next;
            nex->next=pre->next;
            pre->next=nex;
            nex=cur->next;
        }
        pre = cur;
        num-=k;
    }
    return preh->next;
}
