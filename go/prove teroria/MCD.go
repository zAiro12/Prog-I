package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	mcd := MCD(x, y)
	fmt.Println("L'MCD tra", x, "e", y, "vale:", mcd)
}

func MCD(a, b int) int {
	c := 0
	if a < b {
		c = a
	} else {
		c = b
	}
	for a%c != 0 || b%c != 0 {
		c--
	}
	return c
}
