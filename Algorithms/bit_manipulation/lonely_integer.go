package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	var lonelyInt, tmp int
	fmt.Scanf("%d", &lonelyInt)
	for i := 1; i < N; i++ {
		fmt.Scanf("%d", &tmp)
		lonelyInt ^= tmp
	}

	fmt.Println(lonelyInt)
}
