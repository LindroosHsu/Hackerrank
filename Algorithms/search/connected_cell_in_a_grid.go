package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func findConnects(grid [][]int, M, N, r, c int) int {
    answer := 0
    if r < 0 || c < 0 || r >= M || c >= N {
        answer = 0

    } else if grid[r][c] == 1 {
        grid[r][c] = 0
        answer = 1 +
               findConnects(grid, M, N, r-1, c) +
               findConnects(grid, M, N, r+1, c) +
               findConnects(grid, M, N, r, c-1) +
               findConnects(grid, M, N, r, c+1) +
               findConnects(grid, M, N, r-1, c-1) +
               findConnects(grid, M, N, r-1, c+1) +
               findConnects(grid, M, N, r+1, c-1) +
               findConnects(grid, M, N, r+1, c+1)
    }
    return answer
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    M, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    N, _ := strconv.Atoi(scanner.Text())

    grid := make([][]int, M)
    for i := 0; i < M; i++ {
        grid[i] = make([]int, N)
        for j := 0; j < N; j++ {
            scanner.Scan()
            v, _ := strconv.Atoi(scanner.Text())
            grid[i][j] = v
        }
    }

    maxConnects := 0
    for r := 0; r < M; r++ {
        for c := 0; c < N; c++ {
            maxConnects = max(maxConnects, findConnects(grid, M, N, r, c))
        }
    }
    fmt.Println(maxConnects)
}
