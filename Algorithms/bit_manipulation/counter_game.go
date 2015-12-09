package main

import (
    "fmt"
)

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for t := 0; t < T; t++ {
        var N uint64
        fmt.Scanf("%d", &N)
        N -= 1

        var playerCnt uint
        for N != 0 {
            N &= (N-1)
            playerCnt++
        }

        if playerCnt & 1 == 1 {
            fmt.Println("Louise")
        } else {
            fmt.Println("Richard")
        }
    }
}
