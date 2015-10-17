//https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree


/*The tree node has data, left child and right child 
struct node
{
    int data;
    node* left;
    node* right;
};

*/
static int height(node * root)
{
    if(root == NULL){
        return 0;
    }
    
    return 1 + max(height(root->left), height(root->right));
}
  

