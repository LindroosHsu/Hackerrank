#include <stdio.h>

int main() {

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */    
    unsigned long long n;
    scanf("%llu", &n);
    
    int counter_arr[100] = {0};
    
    for(unsigned long long i=0; i<n; i++){
        int tmp_input;
        scanf("%d", &tmp_input);
        counter_arr[tmp_input]++;
    }
    
    for(int i=0, *arr_ptr=counter_arr; i<100; i++, arr_ptr++){
        if(*arr_ptr != 0){
            for(int j=0; j<*arr_ptr; j++){
                printf("%d ", i);
            }
        }
    }

    return 0;
}

