package main

import (
    "fmt"
)

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    times := make(map[int]int)
    // as max overshoot, the answer
    times[0] = 0
    maxD := 0
    for t := 0; t < T; t++ {
        var M, D int
        fmt.Scanf("%d %d\n", &D, &M)

        if D > maxD {
            for i := D; i > maxD; i-- {
                if _, ok := times[i]; !ok {
                    // does not mean anything, but only mach map require
                    times[i] = 1
                    M--

                    if M == 0 {
                        break
                    }
                }
            }
        }

        if M > 0 {
            times[0] += M

            if D > maxD {
                maxD = D
            }
        }

        fmt.Println(times[0])
    }
}
