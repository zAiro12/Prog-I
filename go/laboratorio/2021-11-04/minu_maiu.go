package main

import "fmt"

func main() {
	var s string
	var min, max bool
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			min = true
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			max = true
		}
	}
	if min && max {
		fmt.Println("contiene sia minuscole che maiuscole")
	} else if min {
		fmt.Println("contiene solo minuscole")
	} else if max {
		fmt.Println("contiene solo maiuscole")
	}
}
