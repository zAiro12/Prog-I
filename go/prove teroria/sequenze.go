package main

import "fmt"

func main() {
	var x, c int
	fmt.Scan(&x)
	for x > 1 {
		if x%2 == 0 {
			x /= 2
			c++
		} else {
			x = 3*x + 1
			c++
		}
		fmt.Println(x)

	}
	for i := 0; i <= c; i++ {
		fmt.Print("*")
	}
	fmt.Print(" = ", c, "\n")
}
