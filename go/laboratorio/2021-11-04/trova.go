package main

import "fmt"

func main() {
	var r rune
	var s string
	var f bool
	fmt.Scanf("%c", &r)
	fmt.Scan(&s)

	for i, a := range s {
		if a == r {
			fmt.Println(i)
			f = true
		}
	}
	if !f {
		fmt.Println("-1")
	}
}
