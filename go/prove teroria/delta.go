package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if delta(a, b, c) {
		fmt.Println("è uguale a 0")
	} else {
		fmt.Println("è diverso da 0")
	}
}

func delta(a, b, c int) bool {
	return b*b-4*a*c == 0
}
