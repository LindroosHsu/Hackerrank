package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()

    N, _ := strconv.Atoi(scanner.Text())
    arr := make([]int, N)
    for i := 0; i < N && scanner.Scan(); i++ {
        arr[i], _ = strconv.Atoi(scanner.Text())
    }

    sort.Ints(arr)

    scanner.Scan()
    P, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    Q, _ := strconv.Atoi(scanner.Text())

    var answer, value int
    if P <= arr[0] {
        value = arr[0] - P
        answer = P
    }

    if arr[N-1] <= Q && value < Q - arr[N-1] {
        value = Q - arr[N-1]
        answer = Q
    }

    for i := 0; i < N-1; i++ {
        mid := (arr[i] + arr[i+1]) / 2
        if mid >= P && mid <= Q && value < min(arr[i+1] - mid, mid - arr[i]) {
            value = min(arr[i+1] - mid, mid - arr[i])
            answer = mid

        } else if arr[i] <= P && P <= arr[i+1] && value < min(P - arr[i], arr[i+1] - P) {
            value = min(P - arr[i], arr[i+1] - P)
            answer = mid

        } else if arr[i] <= Q && Q <= arr[i+1] && value < min(Q - arr[i], arr[i+1] - Q) {
            value = min(Q - arr[i], arr[i+1] - Q)
            answer = Q
        }
    }

    fmt.Println(answer)
}
