package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Scan(&in)
	s := []rune(in)
	for i := 0; i < len(s)-1; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z') && (s[i+1] >= '0' && s[i+1] <= '9') {
			fmt.Println(string(s[i]))
		}
	}
}
