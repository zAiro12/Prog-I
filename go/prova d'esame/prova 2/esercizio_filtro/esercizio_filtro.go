package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			fmt.Print(string(input[i]))
		}
	}
	fmt.Println()
}
