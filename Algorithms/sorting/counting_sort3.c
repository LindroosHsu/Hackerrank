#include <stdio.h>

int main() {

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */    
    unsigned long long n;
    scanf("%llu", &n);
    
    int counter_arr[100] = {0};
    
    for(unsigned long long i=0; i<n; i++){
        int tmp_input;
        char tmp_str[11];
        
        scanf("%d %s", &tmp_input, tmp_str);
        counter_arr[tmp_input]++;
    }
    
    int tmp_cache = 0;
    for(int i=0, *arr_ptr=counter_arr; i<100; i++, arr_ptr++){
        *arr_ptr += tmp_cache;
        printf("%d ", *arr_ptr);
        tmp_cache = *arr_ptr;
    }

    return 0;
}
