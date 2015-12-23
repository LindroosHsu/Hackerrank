package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "strconv"
    "container/heap"
)

type pair struct {
    n int
    w int
}

type PriorityQueue []*pair

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].w < pq[j].w }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    T, _ := strconv.Atoi(scanner.Text())
    for t := 0; t < T; t++ {
        scanner.Scan()
        line := strings.Split(scanner.Text(), " ")
        N, _ := strconv.Atoi(line[0])
        M, _ := strconv.Atoi(line[1])

        grid := make(map[int]map[int]int)
        for i := 0; i < M; i++  {
            scanner.Scan()
            line := strings.Split(scanner.Text(), " ")
            x, _ := strconv.Atoi(line[0])
            y, _ := strconv.Atoi(line[1])
            r, _ := strconv.Atoi(line[2])

            if _, ok := grid[x]; !ok {
                grid[x] = make(map[int]int)
            }

            if w, ok := grid[x][y]; !ok || w > r {
                grid[x][y] = r
            }

            if _, ok := grid[y]; !ok {
                grid[y] = make(map[int]int)
            }

            if w, ok := grid[y][x]; !ok || w > r {
                grid[y][x] = r
            }
        }

        scanner.Scan()
        S, _ := strconv.Atoi(scanner.Text())

        dis := make(map[int]int)
        pQ := make(PriorityQueue, 0)
        heap.Init(&pQ)
        heap.Push(&pQ, &pair{S, 0})
        for pQ.Len() > 0 {
            curr := heap.Pop(&pQ).(*pair)

            if _, ok := dis[curr.n]; ok {
                continue
            }
            dis[curr.n] = curr.w

            for des, _ := range grid[curr.n] {
                cost := dis[curr.n] + grid[curr.n][des]
                if _, ok := dis[des]; ok {
                    if dis[des] < cost {
                        continue
                    }
                }
                heap.Push(&pQ, &pair{des, cost})
            }
    	}

        for n := 1; n <= N; n++ {
            if n != S {
                if d, ok := dis[n]; ok {
                    fmt.Printf("%d ", d)
                } else {
                    fmt.Print("-1 ")
                }
            }
        }
        fmt.Println()
    }
}
