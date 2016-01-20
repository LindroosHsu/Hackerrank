// https://www.reddit.com/r/learnprogramming/comments/2ui76s/help_with_solving_a_question_in_logn_time/
//
// will fail on test 5~7
//
package main

import (
	"fmt"
    "bufio"
    "os"
    "strconv"
	"sort"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    N, _ := strconv.Atoi(scanner.Text())

    nArr := make([]int, N)
    for i := 0; i < N && scanner.Scan(); i++ {
        v, _ := strconv.Atoi(scanner.Text())
        nArr[i] = v
    }

    scanner.Scan()
    Q, _ := strconv.Atoi(scanner.Text())

    qArr := make([]int, Q)
    for i := 0; i < N && scanner.Scan(); i++ {
        v, _ := strconv.Atoi(scanner.Text())
        qArr[i] = v
    }

	sort.Ints(nArr)
	cache := make([]int, N + 1)
	for i := 0; i < N; i++ {
		cache[i+1] = cache[i] + nArr[i]
	}

	qSum := 0
	writer := bufio.NewWriter(os.Stdout)
	for _, q := range(qArr) {
		qSum += q
		i := sort.Search(N, func(i int) bool { return nArr[i] >= -qSum })
		big := N - i
		small := i

		bigSum := big * qSum + cache[N] - cache[i]
		smallSum := -cache[i] - small * qSum
		fmt.Fprintln(writer, bigSum + smallSum)
	}
	writer.Flush()
}
