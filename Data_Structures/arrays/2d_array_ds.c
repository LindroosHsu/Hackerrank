#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

#define MAX_ROW 6
#define MAX_COL 6

int main() {

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */  
    int matrix[MAX_ROW][MAX_COL];
    int max_sum = -100;
    
    for(int i = 0; i < MAX_ROW; i++){
        for(int j = 0; j < MAX_COL; j++){
            scanf("%d", &matrix[i][j]);
        }
    }
    
    for(int i = 0; i < MAX_ROW-2; i++){
        for(int j = 0; j < MAX_COL-2; j++){
            
            int tmp_sum =  matrix[i][j] + 
                           matrix[i][j+1] + 
                           matrix[i][j+2] + 
                           matrix[i+1][j+1] + 
                           matrix[i+2][j] + 
                           matrix[i+2][j+1] + 
                           matrix[i+2][j+2];
            
            if(tmp_sum > max_sum){
                max_sum = tmp_sum;
            }
        }
    }
    
    printf("%d\n", max_sum);
    
    return 0;
}

