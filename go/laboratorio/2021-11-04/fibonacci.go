package main

import "fmt"

func main() {
	var x, f int
	fmt.Print("un numero: ")
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		f = fibonacci(i)
		for j := 0; j < f; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func fibonacci(x int) int {
	var a, b int = 0, 1
	for i := 0; i <= x; i++ {
		a, b = b, a+b
	}
	return a
}
