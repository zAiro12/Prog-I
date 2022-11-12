package main

import (
	"fmt"
)

func main() {
	var n, x, c, y int
	fmt.Print("Quanti numeri di mersenne vuoi stampare? ")
	fmt.Scan(&n)

	for x, c = 1, 0; c < n; x *= 2 {
		if isPrime(x - 1) {
			y = x - 1
			c++
		}
	}
	fmt.Println(y)
}

func isPrime(x int) bool {
	for d := 2; d < x; d++ {
		if x%d == 0 {
			return false
		}
	}
	return true
}
