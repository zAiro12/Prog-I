package main

import "fmt"

func main() {
	var min, max int
	fmt.Print("due int: ")
	fmt.Scan(&min, &max)
	if !(max > min) {
		c := max
		max = min
		min = c
	}
	fmt.Println(max)
}
