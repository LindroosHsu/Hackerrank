//https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists

/*
  Merge two sorted lists A and B as one linked list
  Node is defined as 
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
static Node* MergeLists(Node *headA, Node* headB)
{
  // This is a "method-only" submission. 
  // You only need to complete this method 
    
    if(headA == NULL) return headB;
    if(headB == NULL) return headA;

    if(headA->data < headB->data) {
        headA->next = MergeLists(headA->next, headB);
        return headA;
    }else{
        headB->next = MergeLists(headB->next, headA);
        return headB;
    }
}

