// https://www.quora.com/What-is-the-logic-used-in-the-HackerRank-Maximise-Sum-problem
// Instead of using a BST or AVL tree like a lot of people below, it was also fine to just keep it all in a list:
// 1. make prefix sum array mod m,
// 2. sort the prefix sums, but keep each sum with its lowest index in the array,
// 3. start the result at the max of the sum array,
// 4. For each entry of the sorted sums, find the next element in the list that has an original index lower
//    than the entry,or move on if the elements are already worse than the current best
//    compare entry-successor%m to current result, replace if better

package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "sort"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

type pair struct {
    value, index int
}

type lessFunc func(p1, p2 *pair) bool

type multiSorter struct {
	pairs []pair
	less  []lessFunc
}

func (ms *multiSorter) Sort(pairs []pair) {
	ms.pairs = pairs
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.pairs)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.pairs[i], ms.pairs[j] = ms.pairs[j], ms.pairs[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.pairs[i], &ms.pairs[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}

	return ms.less[k](p, q)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()

    T, _ := strconv.Atoi(scanner.Text())
    for T > 0 {
        scanner.Scan()
        N, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        M, _ := strconv.Atoi(scanner.Text())

        prefix := make([]pair, N)
        preSum := 0
        mx, mn := -1, M + 1
        for i := 0; i < N && scanner.Scan(); i++ {
            e, _ := strconv.Atoi(scanner.Text())
            preSum = ((e % M) + preSum) % M
            prefix[i].value = preSum
            prefix[i].index = i
            mx = max(mx, preSum)
        }

        value := func(p1, p2 *pair) bool {
		return p1.value < p2.value
    	}
    	index := func(p1, p2 *pair) bool {
    		return p1.index < p2.index
    	}
        OrderedBy(value, index).Sort(prefix)

        for i := 0; i < N - 1; i++ {
            if prefix[i+1].index < prefix[i].index {
                mn = min(prefix[i+1].value - prefix[i].value, mn)
            }
        }
        fmt.Println(max(mx, (M - mn)))
        T--
    }
}
