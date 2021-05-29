package main

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
)

func main() {
    i := bufio.NewScanner(os.Stdin)
    fmt.Print("Name: ")
    //name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    i.Scan()
    name := i.Text()
    fmt.Print("Address: ")
    i.Scan()
    address := i.Text()
    j, e := json.Marshal(map[string]string{"name": name, "address": address})
    if e == nil {
        fmt.Println(string(j))
    }
}