//https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail

/*
  Get Nth element from the end in a linked list of integers
  Number of elements in the list will always be greater than N.
  Node is defined as 
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
static int GetNode(Node *head,int positionFromTail)
{
  // This is a "method-only" submission. 
  // You only need to complete this method. 
    Node *loop_node = head;
    int len_of_list = 0;
    int counter = 0;
    
    while(loop_node->next != NULL){
        len_of_list++;
        loop_node = loop_node->next;
    }

    loop_node = head;
    while(loop_node->next != NULL){
        if(abs(counter-len_of_list) == positionFromTail){
            break;
        }
        loop_node = loop_node->next;
        counter++;
    }
    
    return loop_node->data;
}

