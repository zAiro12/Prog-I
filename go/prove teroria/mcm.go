package main

import "fmt"

func main() {
	var n1, n2, d1, d2 int
	fmt.Scan(&n1, &d1)
	fmt.Scan(&n2, &d2)
	den := mcm(d1, d2)
	num := (n1*den)/d1 + (n2*den)/d2
	fmt.Println(num, "/", den)
}

func mcm(a, b int) int {
	c := 0
	if a < b {
		c = b
	} else {
		c = a
	}
	for c%a != 0 || c%b != 0 {
		c++
	}
	return c
}
