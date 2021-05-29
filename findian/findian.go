package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("Enter a word: ")
	x, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	x = strings.TrimSpace(x)
	/*
	  fmt.Println(x)
	  fmt.Println(rune(unicode.ToLower(rune(x[0]))))
	  fmt.Println(rune(unicode.ToLower(rune(x[len(x)-1]))))
	  fmt.Println(strings.Contains(strings.ToLower(x[1:len(x)-1]),"a"))
	*/
	if len(x) > 2 && unicode.ToLower(rune(x[0])) == 'i' && unicode.ToLower(rune(x[len(x)-1])) == 'n' && strings.Contains(strings.ToLower(x[1:len(x)-1]), "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
