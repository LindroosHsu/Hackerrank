package main

import (
    "fmt"
)

func isPrime(N int) bool {
    for i := 2; i * i <= N; i++ {
        if N % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for t := 0; t < T; t++ {
        var N int
        fmt.Scanf("%d\n", &N)

        P := make([]int, N+1)
        P[0] = 1
        for i := 1; i <= N; i++ {
            P[i] += P[i-1]
            if i - 4  >= 0 {
                P[i] += P[i-4]
            }
        }

        answer := 0
        for i := 2; i <= P[N]; i++ {
            if isPrime(i) {
                answer++
            }
        }

        fmt.Println(answer)
    }
}
