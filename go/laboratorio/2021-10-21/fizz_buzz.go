package main

import "fmt"

func main() {
	var x int

	fmt.Print("numero: ")
	fmt.Scan(&x)
	if x%3 == 0 && !(x%5 == 0) {
		fmt.Println("Fizz")
	} else if !(x%3 == 0) && (x%5 == 0) {
		fmt.Println("Buzz")
	} else if (x%3 == 0) && (x%5 == 0) {
		fmt.Println("Fizz Buzz")
	}
}
