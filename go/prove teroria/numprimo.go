package main

import (
	"fmt"
	"math"
)

func main() {
	var x, d int
	fmt.Scan(&x)
	for d = 2; d < int(math.Sqrt(float64(x))); d++ {
		if x%d == 0 {
			break
		}
	}
	if d < int(math.Sqrt(float64(x))) {
		fmt.Println(x, "non è primo")
	} else {
		fmt.Println(x, "è primo")
	}
}

func isPrime(x int) bool {
	var d int
	for d = 2; d < int(math.Sqrt(float64(x))); d++ {
		if x%d == 0 {
			break
		}
	}
	if d < int(math.Sqrt(float64(x))) {
		return false
	} else {
		return true
	}
}
