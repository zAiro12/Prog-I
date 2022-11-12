package main

import "fmt"

func main() {
	var x uint
	x = 1
	for {
		x *= 3
		fmt.Println(x)
	}
}
