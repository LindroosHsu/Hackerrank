//https://www.hackerrank.com/challenges/detect-whether-a-linked-list-contains-a-cycle

/*
  Detect loop in a linked list 
  List could be empty also
  Node is defined as 
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
static int HasCycle(Node* head)
{
   // Complete this function
   // Do not write the main method
    
    Node *loop_node = head;
    Node *check_node = head;
    
    while(loop_node != NULL && check_node != NULL){
        loop_node = loop_node->next->next;
        check_node = check_node->next;
        
        if(loop_node == check_node){
            return 1;
        }
    }
    
    return 0;
}

