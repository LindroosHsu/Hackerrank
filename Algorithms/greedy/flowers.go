package main

import (
    "fmt"
    "sort"
)

func main() {
    var N, K int
    fmt.Scanf("%d %d\n", &N, &K)

    flowers := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &flowers[i])
    }

    sort.Ints(flowers)

    var cost, c int
    for i := N-1; i >= 0; i-- {
        cost += ((c/K) + 1) * flowers[i]
        c++
    }

    fmt.Println(cost)
}
