//https://www.quora.com/What-is-the-algorithmic-approach-to-solve-hackerrank-problem-Substring-Diff
// O(N^2) sulotion

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func substringDiff(S, N, pIdx, qIdx int, P, Q string) (maxL int) {
	var mismatches, k int

	// the problem mainained k and l are the same meaning
    for {
        if pIdx + k >= N || qIdx + k >= N {
            maxL = max(maxL, k)
            break
        }

        if P[pIdx + k] != Q[qIdx + k] {
            mismatches++
        }

        if mismatches > S {
            maxL = max(maxL, k)

            for P[pIdx] == Q[qIdx] {
                pIdx++
                qIdx++
                k--
            }
            mismatches--
            pIdx++
            qIdx++

        } else {
            k++
        }
    }

	return maxL
}

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	for t := 0; t < T; t++ {
		var S int
		fmt.Scanf("%d", &S)
		var P, Q string
		fmt.Scanf("%s %s\n", &P, &Q)

        // P and Q has the same length
    	maxL, N := 0, len(P)
        for i := 0; i < N; i++ {
            maxL = max(
                    	max(
                        	maxL,
                        	substringDiff(S, N, 0, i, P, Q)),
                    	substringDiff(S, N, i, 0, P, Q))
        }
		fmt.Println(maxL)
	}
}
