package main

import (
	"fmt"
	"math"
)

func pi(n int64) int64 {
	// Quanti numeri primi ci sono da 1 a n
	var i int64
	var c int64 = 0
	for i = 0; i <= n; i++ {
		if isPrime(i) {
			c++
		}
		fmt.Println("testo il caso", i)
	}
	return c
}

func isPrime(x int64) bool {
	var d int64
	for d = 2; d < int64(math.Sqrt(float64(x))); d++ {
		if x%d == 0 {
			break
		}
	}
	if d < int64(math.Sqrt(float64(x))) {
		return false
	} else {
		return true
	}
}

func main() {
	const num = 10000000
	//9223372036854775807
	//20000000
	fmt.Println("Se il limite per n che tende a infinito del rapporto tra pi(n) e n/log n è 1, allora le sue funzioni sono asintotiche.")
	n1 := float64(pi(num))
	n2 := float64(num) / math.Log(num)
	fmt.Println("Il valore di pi(n) è:", n1)
	fmt.Println("Il valore di n/log n è:", n2)
	fmt.Println("Il valore del loro rapoorto è:", n1/n2)
}
