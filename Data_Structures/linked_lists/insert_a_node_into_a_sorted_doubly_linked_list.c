//https://www.hackerrank.com/challenges/insert-a-node-into-a-sorted-doubly-linked-lists

/*
    Insert Node in a doubly sorted linked list 
    After each insertion, the list should be sorted
   Node is defined as
   struct Node
   {
     int data;
     Node *next;
     Node *prev;
   }
*/
static Node* SortedInsert(Node *head,int data)
{
    // Complete this function
   // Do not write the main method. 
    Node *loop_node = NULL;
    Node *prev_node = NULL;
    Node *new_node = new Node;
    
    new_node->data = data;
    new_node->next = NULL;
    new_node->prev = NULL;
    
    if(head == NULL){
        return new_node;
    }
    
    if(data <= head->data){
        new_node->next = head;
        head->prev = new_node;
        return new_node;
    }
    
    loop_node = head->next;
    prev_node = head;
    while(loop_node != NULL){
        if(data >= prev_node->data && data <= loop_node->data){
            prev_node->next = new_node;
            loop_node->prev = new_node;
            new_node->next = loop_node;
            new_node->prev = prev_node;
                
            return head;
        }
        
        prev_node = loop_node;
        loop_node = loop_node->next;
    }
    
    new_node->prev = prev_node;
    prev_node->next = new_node;
    
    return head;
}

