package main

import (
    "fmt"
)

type node struct {
    x, y, wand int
}

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for T > 0 {
        var N, M int
        fmt.Scanf("%d %d\n", &N, &M)

        grid := make([][]int, N)
        start := node{}
        end := node{}
        for i := 0; i < N; i++ {
            grid[i] = make([]int, M)
            for j := 0; j < M; j++ {
                var e int
                fmt.Scanf("%c", &e)

                if e == 'M' {
                    start.x = j
                    start.y = i

                } else if e == '*' {
                    end.x = j
                    end.y = i
                }

                grid[i][j] = e
            }
            fmt.Scanf("\n")
        }
        var K int
        fmt.Scanf("%d\n", &K)

        currWand := 0
        nextStack := []node{start}
        for len(nextStack) != 0 {
            var ways int
            var left, right, up, down bool

            // pop out a node from stack
            curr := nextStack[len(nextStack)-1]
            nextStack = nextStack[:len(nextStack)-1]

            currWand = curr.wand

            if curr.x == end.x && curr.y == end.y {
                break
            }

            grid[curr.y][curr.x] = 'X'

            if curr.x + 1 < M && grid[curr.y][curr.x+1] != 'X'  {
                right = true
                ways++
            }
            if curr.x > 0 && grid[curr.y][curr.x-1] != 'X' {
                left = true
                ways++
            }
            if curr.y > 0 && grid[curr.y-1][curr.x] != 'X' {
                up = true
                ways++
            }
            if curr.y + 1 < N && grid[curr.y+1][curr.x] != 'X' {
                down = true
                ways++
            }

            if ways > 1 {
                currWand++
            }

            if left {
                nextStack = append(nextStack, node{curr.x-1, curr.y, currWand})
            }
            if right {
                nextStack = append(nextStack, node{curr.x+1, curr.y, currWand})
            }
            if up {
                nextStack = append(nextStack, node{curr.x, curr.y-1, currWand})
            }
            if down {
                nextStack = append(nextStack, node{curr.x, curr.y+1, currWand})
            }
        }

        if currWand == K {
            fmt.Println("Impressed")
        } else {
            fmt.Println("Oops!")
        }

        T--
    }
}
