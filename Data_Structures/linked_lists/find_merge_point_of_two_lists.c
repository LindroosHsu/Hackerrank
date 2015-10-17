//https://www.hackerrank.com/challenges/find-the-merge-point-of-two-joined-linked-lists

/*
   Find merge point of two linked lists
   Node is defined as
   struct Node
   {
       int data;
       Node* next;
   }
*/
static int FindMergeNode(Node *headA, Node *headB)
{
    // Complete this function
    // Do not write the main method. 
    Node *tmp_B = NULL;
    
    while(headA != NULL){
        tmp_B = headB;
        
        while(tmp_B != NULL){
            if(headA == tmp_B){
                return tmp_B->data;
            }
            
            tmp_B = tmp_B->next;
        }
        
        headA = headA->next;
    }
    
    return headA->data;
}

