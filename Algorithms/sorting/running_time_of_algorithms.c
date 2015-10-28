#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>

void static insertion_sort(int arr_size, int *arr) {
    int counter = 0;
    
    for(int i=1; i<arr_size; i++){
        
        int curr = arr[i];
        int pos = i;
        
        while(pos > 0 && arr[pos-1] > curr){
            arr[pos] = arr[pos-1];
            pos--;
            counter++;
        }
        arr[pos] = curr;
    }
    printf("%d\n", counter);
}

int main(void) {
   
    int n;
    scanf("%d", &n);
    
    int arr[n];
    for(int i=0; i<n; i++) { 
        scanf("%d", &arr[i]);
    }

    insertion_sort(n, arr);
   
   return 0;
}
