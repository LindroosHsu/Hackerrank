package main

import (
    "fmt"
    "sort"
)

func checkGrid(grid [][]int, N int, ch chan string) {
    for i := 0; i < N; i++ {
        for j := 1; j < N; j++{
            if grid[j-1][i] > grid[j][i] {
                ch <- "NO"
            }
        }
    }

    ch <- "YES"
}

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for k := 0; k < T; k++ {
        var N, char int
        fmt.Scanf("%d", &N)

        grid := make([][]int, N)
        ch := make(chan string)
        for i := 0; i < N; i++ {
            grid[i] = make([]int, N)
            for j := 0; j < N+1; j++ {
                fmt.Scanf("%c", &char)

                // using this way because I do not want to
                // use scanln then split and covert into int
                if j < N {
                    grid[i][j] = char
                }
            }
            sort.Ints(grid[i])
        }

        go checkGrid(grid, N, ch)

        fmt.Println(<-ch)
    }
}
