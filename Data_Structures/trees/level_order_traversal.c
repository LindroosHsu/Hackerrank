//https://www.hackerrank.com/challenges/tree-level-order-traversal

/*
struct node
{
    int data;
    node* left;
    node* right;
}*/

static void print_level(node *n, int level){
    if(n == NULL){
        return;
    }
    
    if(level == 1){
        cout << n->data << " ";
    }
    else if (level > 1){
        print_level(n->left, level-1);
        print_level(n->right, level-1);
    }
}

static int tree_hight(node *n){
    if(n == NULL){
        return 0;
    }
    
    return 1 + max(tree_hight(n->left), tree_hight(n->right));
}

static void LevelOrder(node * root)
{
    int hight = tree_hight(root);
    int i;
    
    for(i=1; i<=hight; i++){
        print_level(root, i);
    }
}

