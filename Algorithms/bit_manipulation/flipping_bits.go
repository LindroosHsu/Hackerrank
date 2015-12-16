package main

import (
    "fmt"    
)

const (
    maxUint32 = 4294967295
)

func main() {
    var T int
    fmt.Scanf("%d\n", &T)

    for t := 0; t < T; t++ {
        var input uint32
        fmt.Scanf("%d\n", &input)

        fmt.Println(input ^ maxUint32)
    }
}
