#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>

void static print_arr(int arr_size, int *arr){
    while(arr_size--){
        printf("%d ", *arr);
        arr++;
    }
    printf("\n");
}

void static insertion_sort(int arr_size, int *arr) {
    for(int i=1; i<arr_size; i++){
        
        int curr = arr[i];
        int pos = i;
        
        while(pos > 0 && arr[pos-1] > curr){
            arr[pos] = arr[pos-1];
            pos--;
        }
        arr[pos] = curr;
    }
    print_arr(arr_size, arr);
}

int main(void) {
   
    int _ar_size;
    scanf("%d", &_ar_size);
    
    int _ar[_ar_size], _ar_i;
    
    for(_ar_i = 0; _ar_i < _ar_size; _ar_i++) { 
        scanf("%d", &_ar[_ar_i]); 
    }

    insertion_sort(_ar_size, _ar);
   
   return 0;
}

