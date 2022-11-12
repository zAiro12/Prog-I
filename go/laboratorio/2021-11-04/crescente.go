package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, a := range s {
		if a < 'a' || a > 'z' {
			fmt.Println("Stringa sbagliata")
			break
		}
	}
	fmt.Print(string(s[0]))
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			fmt.Print("-")
		}
		fmt.Print(string(s[i+1]))
	}
	fmt.Println()
}
