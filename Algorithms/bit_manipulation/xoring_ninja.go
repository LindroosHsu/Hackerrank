// http://math.stackexchange.com/questions/712487/finding-xor-of-all-subsets
// https://www.quora.com/What-is-the-algorithm-of-this-Hacker-Rank-question
// Consider this:
// XOR_SUM(A,B) = A + B + A^B
// Any 1-bits A and B have in common would be zero-bits in A^B. Any 1-bits they do not have in common would be 1-bits in A^B.
// Put another way, A^B has a 1-bit everywhere A and B are different. Therefore each 1-bit's position is added to the sum two times; once in A or B and once again in A^B.
// Suppose A = 1 and B = 3. Then we have A=0001, B=0011, and x=0011. You can see that A^B = 0010. So
// A + B + A^B = 0b0001 + 0b0010 + 0b0011
//
//   0001 <==   0 +   1
//   0011 <==   2 +   1
// + 0010 <==   2 +   0
// ======
//   0110 <== 2*2 + 2*1
// Each bit appears in the terms exactly two times.
// If you try, you can prove this extends to 3 terms, 4 terms, and n terms, where each bit is added to the sum 2^(n-1) times.

//************************************************
// TEST CASE 6 ~ 13 will get wrong answer...
//************************************************

package main

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
)

const (
    mod int64 = int64(1e9) + 7
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    T, _ := strconv.Atoi(scanner.Text())

    for t := 0; t < T; t++ {
        scanner.Scan()
        N, _ := strconv.ParseUint(scanner.Text(), 10, 64)
        scanner.Scan()
        line := strings.Split(scanner.Text(), " ")

        var xorSum int64
        for _, s := range line {
            v, _ := strconv.ParseInt(s, 10, 64)
            xorSum |= v
        }
        pwr := int64(1 << (N-1))
        fmt.Println((xorSum * pwr) % mod)
    }
}
