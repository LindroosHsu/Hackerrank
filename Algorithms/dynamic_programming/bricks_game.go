package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func min(arr ...int) int {
    m := arr[0]
    for _, a := range arr[1:] {
        if a < m {
            m = a
        }
    }
    return m
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    T, _ := strconv.Atoi(scanner.Text())

    for t := 0; t < T && scanner.Scan(); t++ {
        N, _ := strconv.Atoi(scanner.Text())

        stacks := make([]int, N)
        bricks := make([]int, N+3)
        for i := 0; i < N && scanner.Scan(); i++ {
            stacks[i], _ = strconv.Atoi(scanner.Text())
        }

        sum := 0
        for i := N - 1; i >= 0; i-- {
            sum += stacks[i]
            bricks[i] = sum - min(bricks[i+1], bricks[i+2], bricks[i+3])
        }

        fmt.Println(bricks[0])
    }
}
