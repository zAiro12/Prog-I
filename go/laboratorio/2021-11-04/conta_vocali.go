package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	var c int
	fmt.Scan(&s)
	for _, a := range s {
		a = unicode.ToLower(a)
		if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' {
			c++
		}
	}
	fmt.Println(c)
}
