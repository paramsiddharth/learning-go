package main

import "fmt"

func main() {
  fmt.Printf("Enter a floating-point value: ")
  var x float64
  fmt.Scan(&x)
  fmt.Printf("The truncated integer is %d.\n",int(x))
}