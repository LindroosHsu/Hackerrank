package main

import (
    "fmt"
    "sort"
)

const (
    freeRange = 4
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    weights := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &weights[i])
    }

    sort.Ints(weights)

    cost := 0
    free := 0
    for i := 0; i < N; i += free+1 {
        free = 0
        for j := i+1; j < N; j++ {
            if weights[j] <= weights[i] + freeRange {
                free++
            } else {
                break
            }
        }
        cost++
    }

    fmt.Println(cost)
}
