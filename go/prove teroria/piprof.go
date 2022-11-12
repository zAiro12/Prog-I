package main

import (
	"fmt"
	"math"
)

func main() {
	for n := 1; ; n++ {
		pin := 0
		if isPrime(int64(n)) {
			pin++
		}
		ex := float64(n) / math.Log(float64(n))
		act := pin
		fmt.Print(n, "\t", act, "\t", ex, "\t", float64(act))
		fmt.Println()
	}
}

func isPrime(x int64) bool {
	var d int64
	for d = 2; d < int64(math.Sqrt(float64(x))); d++ {
		if x%d == 0 {
			break
		}
	}
	return d == int64(math.Sqrt(float64(x)))
}

func pi(n int64) int {
	c := 0
	var i int64
	for i = 1; i <= n; i++ {
		if isPrime(i) {
			c++
		}
	}
	return c
}
