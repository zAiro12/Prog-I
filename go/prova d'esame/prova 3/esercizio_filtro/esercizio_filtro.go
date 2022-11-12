package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ch := os.Args[1:]
	for i := 0; i < len(ch); i += 2 {
		n, _ := strconv.Atoi(ch[i+1])
		for j := 0; j < n; j++ {
			fmt.Print(ch[i])
		}
	}
	fmt.Println()
}
