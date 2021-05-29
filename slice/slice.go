package main

import (
    "fmt"
    "strconv"
    "sort"
)

func main() {
    var s []int = make([]int, 0, 3)
    var x string
    fmt.Scan(&x)
    
    for x != "X" {
        n, e := strconv.Atoi(x)
        if e == nil {
            s = append(s, n)
        } else {
            fmt.Println("Enter a valid integer.")
        }
        sort.Ints(s)
        for _, v := range s {
            fmt.Printf("%d ", v)
        }
        fmt.Println()
        fmt.Scan(&x)
    }
}