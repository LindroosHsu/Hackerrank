#include <stdio.h>

// the global variable for counting
static int m_shift = 0;
static int m_swap = 0;

static void swap(int *x, int *y){
    *x=*x^*y^(*y=*x);
}

static int partition(int *arr, int lo, int hi){
    int pivot = arr[hi];
    int swp_pos = lo;
    
    for(int i=lo; i<hi; i++){
        if(arr[i] <= pivot){
            swap(&arr[swp_pos], &arr[i]);
            swp_pos++;
            m_swap++;
        }
    }
    
    swap(&arr[swp_pos], &arr[hi]);
    m_swap++;
    
    return swp_pos;
}

static void quick_sort(int *arr, int lo, int hi){
    if(lo < hi){
        // partition index
        int p = partition(arr, lo, hi);
        
        quick_sort(arr, lo, p-1);
        quick_sort(arr, p+1, hi);
    }
}

static void insertion_sort(int *arr, int arr_size) {
    for(int i=1; i<arr_size; i++){
        
        int curr = arr[i];
        int pos = i;
        
        while(pos > 0 && arr[pos-1] > curr){
            arr[pos] = arr[pos-1];
            pos--;
            m_shift++;
        }
        arr[pos] = curr;
    }
}


int main(void) {

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */  
    int n;
    scanf("%d", &n);
    
    int q_arr[n];
    int i_arr[n];
    int *q_arr_ptr = q_arr;
    int *i_arr_ptr = i_arr;
    
    for(int i=0; i<n; i++, q_arr_ptr++, i_arr_ptr++){
        int tmp_input;
        scanf("%d", &tmp_input);
        *q_arr_ptr = tmp_input;
        *i_arr_ptr = tmp_input;
    }
    
    quick_sort(q_arr, 0, n-1);
    insertion_sort(i_arr, n);
    
    printf("%d\n", m_shift-m_swap);
    return 0;
}

