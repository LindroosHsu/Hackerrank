//https://www.hackerrank.com/challenges/tree-top-view

/*
struct node
{
    int data;
    node* left;
    node* right;
};

*/
static void left_side(node *root){
    if(root == NULL){
        return;
    }
    
    left_side(root->left);
    cout << root->data << " ";
}

static void right_side(node *root){
    if(root == NULL){
        return;
    }
    
    cout << root->data << " ";
    right_side(root->right);
}

static void top_view(node * root)
{
    if(root == NULL){
        return;
    }
    
    left_side(root->left);
    cout << root->data << " ";
    right_side(root->right);
}

