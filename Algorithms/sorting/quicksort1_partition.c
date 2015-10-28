#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>

void static print_arr(int arr_size, int *arr){
    for(int i=0; i<arr_size; i++){
        printf("%d ", *arr);
        arr++;
    }
    printf("\n");
}

void static partition(int arr_size, int *arr) {
    int p = arr[0];
    int answer[arr_size];
    
    // using pointer way to store value
    int *pos_right = &answer[arr_size-1];
    int *pos_left = answer;
    
    // starting from arr[1]
    arr++;
    for(int i=1; i<arr_size; i++){
        if(*arr < p){
            *pos_left = *arr;
            pos_left++;
        }
        else if(*arr > p){
            *pos_right = *arr;
            pos_right--;
        }
        arr++;
    }
    
    // all setting done, insert the p value
    *pos_left = p;
    
    print_arr(arr_size, answer);
}

int main(void) {
   
    int _ar_size;
    scanf("%d", &_ar_size);
    
    int _ar[_ar_size], _ar_i;
    for(_ar_i = 0; _ar_i < _ar_size; _ar_i++) { 
        scanf("%d", &_ar[_ar_i]); 
    }

    partition(_ar_size, _ar);
   
   return 0;
}

