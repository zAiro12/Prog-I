package main

import "fmt"

func main() {
	var totLen, c int
	var s string
	fmt.Scan(&totLen)
	for c = 0; c < totLen; {
		fmt.Scan(&s)
		c += len(s)
	}
	fmt.Println(c)
}
