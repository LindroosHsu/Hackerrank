#include <stdio.h>

void static swap(int *x, int *y){
    *x=*x^*y^(*y=*x);
}

void static print_arr(int *arr, int n){
    while(n--){
        printf("%d ", *arr);
        arr++;
    }
    printf("\n");
}

int static partition(int *arr, int lo, int hi){
    int pivot = arr[hi];
    int swp_pos = lo;
    
    for(int i=lo; i<hi; i++){
        if(arr[i] <= pivot){
            swap(&arr[swp_pos], &arr[i]);
            swp_pos++;
        }
    }
    
    swap(&arr[swp_pos], &arr[hi]);
    return swp_pos;
}

void static quick_sort(int *arr, int n, int lo, int hi){
    if(lo < hi){
        // partition index
        int p = partition(arr, lo, hi);
        print_arr(arr, n);
        
        quick_sort(arr, n, lo, p-1);
        quick_sort(arr, n, p+1, hi);
    }
}

int main() {

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */  
    int n;
    scanf("%d", &n);
    
    int arr[n];
    int *arr_ptr = arr;
    
    for(int i=0; i<n; i++, arr_ptr++){
        scanf("%d", arr_ptr);
    }
    
    quick_sort(arr, n, 0, n-1);
    
    return 0;
}
