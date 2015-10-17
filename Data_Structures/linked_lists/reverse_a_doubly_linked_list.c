//https://www.hackerrank.com/challenges/reverse-a-doubly-linked-lists

/*
   Reverse a doubly linked list, input list may also be empty
   Node is defined as
   struct Node
   {
     int data;
     Node *next;
     Node *prev;
   }
*/
static Node* Reverse(Node* head)
{
    // Complete this function
    // Do not write the main method. 
    Node *curr_node = head;
    Node *prev_node = NULL;
    Node *next_node = NULL;
    
    if(head == NULL){
        return head;
    }
    
    while(1){
        prev_node = curr_node->prev;
        next_node = curr_node->next;
        
        curr_node->next = prev_node;
        curr_node->prev = next_node;
        
        if(next_node == NULL){
            break;
        }
        
        curr_node = next_node;
    }
    
    return curr_node;
}

