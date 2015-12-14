// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

package main

import (
    "fmt"
    "bufio"
    "strconv"
    "os"
    "strings"
    "container/heap"
)

type node struct {
    num int
    fare int
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].fare < pq[j].fare }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    // skip E
    scanner.Scan()
    line := strings.Split(scanner.Text(), " ")
    N, _ := strconv.Atoi(line[0])

    grid := make(map[int][]*node)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        A, _ := strconv.Atoi(line[0])
        B, _ := strconv.Atoi(line[1])
        C, _ := strconv.Atoi(line[2])

        grid[A] = append(grid[A], &node{B, C})
        // bidirectional roads
        grid[B] = append(grid[B], &node{A, C})
    }

    answer := -1
    costCache := make(map[int]int)
    stationQueue := make(PriorityQueue, 0)
    heap.Init(&stationQueue)
    heap.Push(&stationQueue, &node{1, 0})
    for stationQueue.Len() > 0 {
        currStation := heap.Pop(&stationQueue).(*node)
		if currStation.num == N {
            answer = currStation.fare
            break
        }

        if _, inCache := costCache[currStation.num]; inCache {
            continue
        }
        costCache[currStation.num] = currStation.fare
        for _, destStation := range grid[currStation.num] {
            baseCost := costCache[currStation.num]
            cost := baseCost + max(0, destStation.fare - baseCost)
            if _, inCache := costCache[destStation.num]; inCache {
                if costCache[destStation.num] < cost {
                    continue
                }
            }
            heap.Push(&stationQueue, &node{destStation.num, cost})
        }
	}

    if answer != -1 {
        fmt.Println(answer)
    } else {
        fmt.Println("NO PATH EXISTS")
    }
}
