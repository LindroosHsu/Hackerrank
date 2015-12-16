package main

import (
    "fmt"
)

func main() {
    var R, L int
    fmt.Scanf("%d\n%d\n", &L, &R)

    answer := 0
    for l := L; l <= R; l++ {
        for r := l; r <= R; r++ {
            tmp := r ^ l
            if tmp > answer {
                answer = tmp
            }
        }
    }

    fmt.Println(answer)
}
