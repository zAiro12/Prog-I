package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(x int) bool {
	var d int
	for d = 2; d <= int(math.Sqrt(float64(x))); d++ {
		if x%d == 0 {
			break
		}
	}
	if d <= int(math.Sqrt(float64(x))) {
		return false
	} else {
		return true
	}
}
