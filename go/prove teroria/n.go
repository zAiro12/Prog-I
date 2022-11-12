package main

import "fmt"

func main() {
	var n, i, r int
	fmt.Scan(&n)
	for i != n {
		r += n - i
		i++
	}
	fmt.Println(r)
}
