#include<stdio.h>

// function to merge the two haves arr[l..m] and arr[m+1..r] of array arr[]
inline static void merge(unsigned int *arr, unsigned int l_idx, unsigned int mid, unsigned int r_idx, long long *counter)
{
    unsigned int i, j, k;
    unsigned int L_size = mid - l_idx + 1;
    unsigned int R_size = r_idx - mid;

    // create temp arrays
    unsigned int L[L_size], R[R_size];

    // copy data to temp arrays L[] and R[]
    for(i=0; i<L_size; i++){
        L[i] = arr[l_idx + i];
    }

    for(j=0; j<R_size; j++){
        R[j] = arr[mid + 1+ j];
    }

    // merge the temp arrays back into arr[l..r]
    i = 0;
    j = 0;
    k = l_idx;
    while (i < L_size && j < R_size){
        if (L[i] <= R[j]){
            arr[k++] = L[i++];
        }
        else{
            arr[k++] = R[j++];
            *counter += L_size - i;
        }
    }

    // Copy the remaining elements of L[], if there are any
    while (i < L_size){
        arr[k++] = L[i++];
    }

    // Copy the remaining elements of R[], if there are any
    while (j < R_size){
        arr[k++] = R[j++];
    }
}

// l_idx is for left index and r_idx is right index of the sub-array of arr to be sorted
inline static void merge_sort(unsigned int *arr, unsigned int l_idx, unsigned int r_idx, long long *counter)
{
    if (l_idx < r_idx)
    {
        //Same as (l+r_idx)/2, but avoids overflow for large l and h
        unsigned int mid = l_idx + (r_idx - l_idx) / 2;

        merge_sort(arr, l_idx, mid, counter);
        merge_sort(arr, mid+1, r_idx, counter);
        merge(arr, l_idx, mid, r_idx, counter);
    }
}

int main(){
    int t;
    scanf("%d", &t);

    while(t--){
        unsigned int n;
        scanf("%d", &n);

        unsigned int arr[n];
        for(unsigned int i=0; i<n; i++){
            scanf("%d", &arr[i]);
        }

        long long counter = 0;
        merge_sort(arr, 0, n-1, &counter);
        printf("%lld\n", counter);
    }

    return 0;
}
