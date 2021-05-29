package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
    "strconv"
)

// Swap values at positions i and i + 1 of slice sli
func Swap(sli []int, i int) {
    temp := sli[i]
    sli[i] = sli[i + 1]
    sli[i + 1] = temp
}

// Sort the slice by bubble sorting algorithm
func BubbleSort(sli []int) {
    for i := 0; i < len(sli) - 1; i++ {
        swaps := 0
        
        for j := 0; j < len(sli) - 1 - i; j++ {
            if sli[j + 1] < sli[j] {
                Swap(sli, j)
                swaps++
            }
        }
        
        // Check if sorted
        if swaps == 0 {
            break
        }
    }
}

func main() {
    // Read input using bufio
    fmt.Println("Sequence of atmost 10 integers:")
    seq, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    seq = strings.TrimSpace(seq)
    
    // Initialize index variable, split string by space and initialize result slice
    i := 0
    seqsplit := strings.Split(seq, " ")
    var toSort []int = make([]int, 0, 10)
    
    // Add the numbers to the slice if valid
    for i < len(seqsplit) && i < 10 {
        n, e := strconv.Atoi(seqsplit[i])
        i++
        if e != nil {
            continue
        }
        toSort = append(toSort, n)
        // fmt.Println(toSort)
    }
    
    // Sort the slice
    BubbleSort(toSort)
    
    // Print the slice
    for _, v := range toSort {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}