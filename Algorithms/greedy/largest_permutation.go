package main

import (
    "fmt"
)

func main() {
    var N, K int
    fmt.Scanf("%d %d\n", &N, &K)

    arr := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &arr[i])
    }

    for i := 0; i < N-1; i++ {
        if K <= 0 {
            break
        }

        maxValue, maxLoc := arr[i], i
        for j := i+1; j < N; j++ {
            if arr[j] > maxValue {
                maxValue = arr[j]
                maxLoc = j
            }
        }

        if maxLoc != i {
            K--
            arr[i], arr[maxLoc] = arr[maxLoc], arr[i]
        }
    }

    for i := 0; i < N; i++ {
        fmt.Printf("%d ", arr[i])
    }
}
