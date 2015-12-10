package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

const (
    maxChocol   = 1001
    maxUint     = ^uint(0)
    maxInt      = int(maxUint >> 1)
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    T, _ := strconv.Atoi(scanner.Text())

    for t := 0; t < T; t++ {
        scanner.Scan()
        N, _ := strconv.Atoi(scanner.Text())

        scanner.Scan()
        line := strings.Split(scanner.Text(), " ")

        var chocols []int
        minChocol := maxChocol
        for _, v := range line {
            inChocol, _ := strconv.Atoi(v)

            if minChocol > inChocol {
                minChocol = inChocol
            }

            chocols = append(chocols, inChocol)
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
