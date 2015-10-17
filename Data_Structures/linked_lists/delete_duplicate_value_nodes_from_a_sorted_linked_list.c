//https://www.hackerrank.com/challenges/delete-duplicate-value-nodes-from-a-sorted-linked-lists

/*
  Remove all duplicate elements from a sorted linked list
  Node is defined as 
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
static Node* RemoveDuplicates(Node *head)
{
  // This is a "method-only" submission. 
  // You only need to complete this method. 
    Node *loop_node1 = head;
    Node *loop_node2 = NULL;
    Node *tmp_node = NULL;
    
    if(head == NULL){
        return NULL;
    }
    
    while(loop_node1 != NULL){
        loop_node2 = loop_node1->next;
        
        while(loop_node2 != NULL && loop_node2->data == loop_node1->data){
            tmp_node = loop_node2;
            loop_node1->next = tmp_node->next;
            loop_node2 = tmp_node->next;
                
            delete tmp_node;
        }
        
        loop_node1 = loop_node1->next;
    }
    
    return head;
}

