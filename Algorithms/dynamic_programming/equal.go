package main

import (
    "fmt"
)

const (
    maxChocol   = 1001
    maxUint     = ^uint(0)
    maxInt      = int(maxUint >> 1)
)

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for t := 0; t < T; t++ {
        var N int
        fmt.Scanf("%d", &N)

        chocols := make([]int, N)
        minChocol := maxChocol
        for i := 0; i < N; i++ {
            var inChocol int
            fmt.Scanf("%d", &inChocol)

            if minChocol > inChocol {
                minChocol = inChocol
            }

            chocols[i] = inChocol
        }

        answer := maxInt
        for i := 0; i <= 5; i++ {
            var counter int
            for j := 0; j < N; j++ {
                v := chocols[j] - (minChocol - i)
                counter += v/5 + v%5/2 + v%5%2
            }

            if counter < answer {
                answer = counter
            }
        }

        fmt.Println(answer)
    }
}
