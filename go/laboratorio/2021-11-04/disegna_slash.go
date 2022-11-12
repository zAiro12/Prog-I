package main

import "fmt"

func main() {
	var x, i int
	fmt.Print("dimensione \\: ")
	fmt.Scan(&x)
	for i = 0; i < x; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println()
	}

}
