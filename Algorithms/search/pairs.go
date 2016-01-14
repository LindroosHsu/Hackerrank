// here's what you can do: Hash the list of integers in the first stage.
// Now, cycle through the list of integers again,
// and for every integer I, look for a corresponding integer K -I in the Hashmap.
// This should do it in approximately linear time. unlike your current code which will require approx quadratic time

package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    N, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    K, _ := strconv.Atoi(scanner.Text())

    numbers := make(map[int]int)
    for i := 0; i < N && scanner.Scan(); i++ {
        v, _ := strconv.Atoi(scanner.Text())
        numbers[v] = 0
    }

    answer := 0
    for key, _ := range numbers {
        if _, ok := numbers[key-K]; ok {
            answer++
        }
    }

    fmt.Println(answer)
}
